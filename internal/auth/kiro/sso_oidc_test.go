package kiro

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

type recordingRoundTripper struct {
	lastReq *http.Request
}

func (rt *recordingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	rt.lastReq = req
	body := `{"nextToken":null,"profiles":[{"arn":"arn:aws:codewhisperer:us-east-1:123456789012:profile/ABC","profileName":"test"}]}`
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     make(http.Header),
	}, nil
}

func TestTryListAvailableProfiles_UsesClientIDForAccountKey(t *testing.T) {
	rt := &recordingRoundTripper{}
	client := &SSOOIDCClient{
		httpClient: &http.Client{Transport: rt},
	}

	profileArn := client.tryListAvailableProfiles(context.Background(), "access-token", "client-id-123", "refresh-token-456")
	if profileArn == "" {
		t.Fatal("expected profileArn, got empty result")
	}

	accountKey := GetAccountKey("client-id-123", "refresh-token-456")
	fp := GlobalFingerprintManager().GetFingerprint(accountKey)
	expected := fmt.Sprintf("aws-sdk-js/%s KiroIDE-%s-%s", fp.RuntimeSDKVersion, fp.KiroVersion, fp.KiroHash)
	got := rt.lastReq.Header.Get("X-Amz-User-Agent")
	if got != expected {
		t.Errorf("X-Amz-User-Agent = %q, want %q", got, expected)
	}
}

func TestTryListAvailableProfiles_UsesRefreshTokenWhenClientIDMissing(t *testing.T) {
	rt := &recordingRoundTripper{}
	client := &SSOOIDCClient{
		httpClient: &http.Client{Transport: rt},
	}

	profileArn := client.tryListAvailableProfiles(context.Background(), "access-token", "", "refresh-token-789")
	if profileArn == "" {
		t.Fatal("expected profileArn, got empty result")
	}

	accountKey := GetAccountKey("", "refresh-token-789")
	fp := GlobalFingerprintManager().GetFingerprint(accountKey)
	expected := fmt.Sprintf("aws-sdk-js/%s KiroIDE-%s-%s", fp.RuntimeSDKVersion, fp.KiroVersion, fp.KiroHash)
	got := rt.lastReq.Header.Get("X-Amz-User-Agent")
	if got != expected {
		t.Errorf("X-Amz-User-Agent = %q, want %q", got, expected)
	}
}

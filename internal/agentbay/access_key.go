// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package agentbay

import (
	"fmt"
	"net/http"

	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"

	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
)

func newSDKClientWithAccessKeys(apiConfig *config.APIConfig, accessKeyID, accessKeySecret, securityToken string) (*client.Client, error) {
	endpoint := apiConfig.Endpoint
	openapiConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyID),
		AccessKeySecret: dara.String(accessKeySecret),
		Endpoint:        dara.String(endpoint),
		ReadTimeout:     dara.Int(apiConfig.TimeoutMs),
		ConnectTimeout:  dara.Int(apiConfig.TimeoutMs),
		UserAgent:       dara.String("AgentBay-CLI/1.0"),
	}
	if securityToken != "" {
		openapiConfig.SecurityToken = dara.String(securityToken)
	}

	baseTransport := http.DefaultTransport
	if baseTransport == nil {
		baseTransport = &http.Transport{}
	}
	debugTransport := &debugTransport{base: baseTransport}
	httpClient := &http.Client{Transport: debugTransport}
	openapiConfig.HttpClient = &debugHttpClient{client: httpClient}

	sdkClient, err := client.NewClient(openapiConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create API client: %w", err)
	}
	return sdkClient, nil
}

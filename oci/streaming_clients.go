// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_streaming "github.com/oracle/oci-go-sdk/v31/streaming"

	oci_common "github.com/oracle/oci-go-sdk/v31/common"
)

func init() {
	RegisterOracleClient("oci_streaming.StreamAdminClient", &OracleClient{initClientFn: initStreamingStreamAdminClient})
}

func initStreamingStreamAdminClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_streaming.NewStreamAdminClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) streamAdminClient() *oci_streaming.StreamAdminClient {
	return m.GetClient("oci_streaming.StreamAdminClient").(*oci_streaming.StreamAdminClient)
}

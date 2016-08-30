/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"bytes"
	"encoding/json"
)

type report struct {
	Code     int    `json:"return_code"`
	Instance string `json:"instance"`
	StdErr   string `json:"stderr"`
	StdOut   string `json:"stdout"`
}

type ServiceOptions struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type builderEvent struct {
	Uuid                  string   `json:"_uuid"`
	BatchID               string   `json:"_batch_id"`
	Type                  string   `json:"type"`
	Service               string   `json:"service"`
	Created               bool     `json:"created"`
	Name                  string   `json:"name"`
	Payload               string   `json:"payload"`
	Target                string   `json:"target"`
	Reports               []report `json:"reports,omitempty"`
	MatchedInstances      []string `json:"matched_instances,omitempty"`
	ServiceName           string   `json:"service_name"`
	ServiceType           string   `json:"service_type"`
	EndPoint              string   `json:"service_endpoint"`
	User                  string   `json:"user"`
	Password              string   `json:"password"`
	RouterIP              string   `json:"router_ip"`
	ClientName            string   `json:"client_name"`
	DatacenterName        string   `json:"datacenter_name"`
	DatacenterPassword    string   `json:"datacenter_password"`
	DatacenterRegion      string   `json:"datacenter_region"`
	DatacenterType        string   `json:"datacenter_type"`
	DatacenterUsername    string   `json:"datacenter_username"`
	DatacenterAccessToken string   `json:"datacenter_access_token"`
	DatacenterAccessKey   string   `json:"datacenter_access_key"`
	NetworkName           string   `json:"network_name"`
	VCloudURL             string   `json:"vcloud_url"`
	Status                string   `json:"status"`
	ErrorCode             string   `json:"error_code"`
	ErrorMessage          string   `json:"error_message"`
}

type vcloudEvent struct {
	Uuid    string `json:"_uuid"`
	BatchID string `json:"_batch_id"`
	Type    string `json:"_type"`

	Created          bool           `json:"created"`
	Name             string         `json:"execution_name"`
	Service          string         `json:"service_id"`
	ServiceName      string         `json:"service_name"`
	ServiceType      string         `json:"service_type"`
	ServiceEndPoint  string         `json:"service_endpoint"`
	ServiceOptions   ServiceOptions `json:"service_options"`
	ExecutionType    string         `json:"execution_type"`
	ExecutionPayload string         `json:"execution_payload"`
	ExecutionTarget  string         `json:"execution_target"`
	ExecutionResults struct {
		Reports []report `json:"reports,omitempty"`
	} `json:"execution_results,omitempty"`
	ExecutionMatchedInstances []string `json:"execution_matched_instances,omitempty"`
	ExecutionStatus           string   `json:"execution_status,omitempty"`
	ClientName                string   `json:"client_name"`
	DatacenterName            string   `json:"datacenter_name"`
	DatacenterPassword        string   `json:"datacenter_password"`
	DatacenterRegion          string   `json:"datacenter_region"`
	DatacenterType            string   `json:"datacenter_type"`
	DatacenterUsername        string   `json:"datacenter_username"`
	NetworkName               string   `json:"network_name"`
	VCloudURL                 string   `json:"vcloud_url"`
	Status                    string   `json:"status"`
	ErrorCode                 string   `json:"error_code"`
	ErrorMessage              string   `json:"error_message"`
}

type Translator struct{}

func (t Translator) BuilderToConnector(j []byte) []byte {
	var input builderEvent
	var output []byte
	json.Unmarshal(j, &input)

	switch input.Type {
	case "salt", "fake", "vcloud-fake":
		output = t.builderToVCloudConnector(input)
	}

	return output
}

func (t Translator) builderToVCloudConnector(input builderEvent) []byte {
	var output vcloudEvent

	output.Uuid = input.Uuid
	output.BatchID = input.BatchID
	output.Service = input.Service
	if input.DatacenterType == "vcloud-fake" || input.DatacenterType == "aws-fake" || input.DatacenterType == "fake" {
		output.Type = "fake"
	} else {
		output.Type = input.Type
	}

	output.Name = input.Name
	output.ServiceName = input.ServiceName
	output.ServiceType = input.ServiceType
	output.ServiceEndPoint = input.EndPoint
	output.ServiceOptions.User = input.User
	output.ServiceOptions.Password = input.Password
	output.ExecutionType = input.Type
	output.ExecutionPayload = input.Payload
	output.ExecutionTarget = input.Target

	output.NetworkName = input.NetworkName
	output.ClientName = input.ClientName
	output.DatacenterName = input.DatacenterName
	output.DatacenterRegion = input.DatacenterRegion
	output.DatacenterUsername = input.DatacenterUsername
	output.DatacenterPassword = input.DatacenterPassword
	output.DatacenterType = input.DatacenterType
	output.VCloudURL = input.VCloudURL
	output.Status = input.Status
	output.ErrorCode = input.ErrorCode
	output.ErrorMessage = input.ErrorMessage

	body, _ := json.Marshal(output)

	return body
}

func (t Translator) ConnectorToBuilder(j []byte) []byte {
	var output []byte
	var input map[string]interface{}

	dec := json.NewDecoder(bytes.NewReader(j))
	dec.Decode(&input)

	switch input["_type"] {
	case "salt", "fake", "vcloud-fake":
		output = t.vcloudConnectorToBuilder(j)
	}

	return output
}

func (t Translator) vcloudConnectorToBuilder(j []byte) []byte {
	var input vcloudEvent
	var output builderEvent
	json.Unmarshal(j, &input)

	output.Uuid = input.Uuid
	output.BatchID = input.BatchID
	output.Service = input.Service
	output.Type = input.DatacenterType

	output.Name = input.Name
	output.ServiceName = input.ServiceName
	output.ServiceType = input.ServiceType
	output.EndPoint = input.ServiceEndPoint
	output.User = input.ServiceOptions.User
	output.Password = input.ServiceOptions.Password
	output.Payload = input.ExecutionPayload
	output.Target = input.ExecutionTarget

	output.NetworkName = input.NetworkName
	output.ClientName = input.ClientName
	output.DatacenterName = input.DatacenterName
	output.DatacenterRegion = input.DatacenterRegion
	output.DatacenterUsername = input.DatacenterUsername
	output.DatacenterPassword = input.DatacenterPassword
	output.DatacenterType = input.DatacenterType
	output.VCloudURL = input.VCloudURL
	output.Status = input.Status
	output.ErrorCode = input.ErrorCode
	output.ErrorMessage = input.ErrorMessage

	body, _ := json.Marshal(output)

	return body
}

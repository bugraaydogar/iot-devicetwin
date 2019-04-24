// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * This file is part of the IoT Device Twin Service
 * Copyright 2019 Canonical Ltd.
 *
 * This program is free software: you can redistribute it and/or modify it
 * under the terms of the GNU Affero General Public License version 3, as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranties of MERCHANTABILITY,
 * SATISFACTORY QUALITY, or FITNESS FOR A PARTICULAR PURPOSE.
 * See the GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package web

import (
	"encoding/json"
	"github.com/CanonicalLtd/iot-devicetwin/domain"
	"log"
	"net/http"
)

// JSONHeader is the header for JSON responses
const JSONHeader = "application/json; charset=UTF-8"

// StandardResponse is the JSON response from an API method, indicating success or failure.
type StandardResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// SnapsResponse is the JSON response to list snaps
type SnapsResponse struct {
	StandardResponse
	Snaps []domain.DeviceSnap `json:"snaps"`
}

// DeviceResponse is the JSON response to list snaps
type DeviceResponse struct {
	StandardResponse
	Device domain.Device `json:"device"`
}

// formatStandardResponse returns a JSON response from an API method, indicating success or failure
func formatStandardResponse(code, message string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := StandardResponse{Code: code, Message: message}

	if len(code) > 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	// Encode the response as JSON
	encodeResponse(w, response)
}

// formatSnapsResponse returns a JSON response from a snap list API method
func formatSnapsResponse(snaps []domain.DeviceSnap, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := SnapsResponse{StandardResponse{}, snaps}

	// Encode the response as JSON
	encodeResponse(w, response)
}

// formatDeviceResponse returns a JSON response from a snap list API method
func formatDeviceResponse(device domain.Device, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := DeviceResponse{StandardResponse{}, device}

	// Encode the response as JSON
	encodeResponse(w, response)
}

func encodeResponse(w http.ResponseWriter, response interface{}) {
	// Encode the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error forming the response:", err)
	}
}

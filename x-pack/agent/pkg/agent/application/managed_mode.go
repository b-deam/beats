// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package application

import (
	"github.com/pkg/errors"

	"github.com/elastic/beats/x-pack/agent/pkg/config"
	"github.com/elastic/beats/x-pack/agent/pkg/core/logger"
)

// Errors generated by the application.
var (
	ErrManagedIsNotImplemented = errors.New("fleet managed mode is not implemented in the agent")
)

// Managed application, when the application is run in managed mode, most of the configuration are
// coming from the Fleet App.
type Managed struct {
	log    *logger.Logger
	client sender
}

func newManaged(
	log *logger.Logger,
	rawConfig *config.Config,
	router *router,
	client sender,
) (*Managed, error) {
	return nil, ErrManagedIsNotImplemented
}

// Start starts a managed agent.
func (m *Managed) Start() error {
	return ErrManagedIsNotImplemented
}

// Stop stops a managed agent.
func (m *Managed) Stop() error {
	return ErrManagedIsNotImplemented
}
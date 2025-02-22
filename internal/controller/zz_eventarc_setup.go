/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	channel "github.com/upbound/provider-gcp/internal/controller/eventarc/channel"
	googlechannelconfig "github.com/upbound/provider-gcp/internal/controller/eventarc/googlechannelconfig"
	trigger "github.com/upbound/provider-gcp/internal/controller/eventarc/trigger"
)

// Setup_eventarc creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_eventarc(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		channel.Setup,
		googlechannelconfig.Setup,
		trigger.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

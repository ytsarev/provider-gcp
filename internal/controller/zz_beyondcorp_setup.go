/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	appconnection "github.com/upbound/provider-gcp/internal/controller/beyondcorp/appconnection"
	appconnector "github.com/upbound/provider-gcp/internal/controller/beyondcorp/appconnector"
	appgateway "github.com/upbound/provider-gcp/internal/controller/beyondcorp/appgateway"
)

// Setup_beyondcorp creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_beyondcorp(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appconnection.Setup,
		appconnector.Setup,
		appgateway.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

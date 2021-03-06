/*
 * ytelapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */

package email_pkg

import "ytelapi_lib/configuration_pkg"

/*
 * Interface for the EMAIL_IMPL
 */
type EMAIL interface {
    CreateEmailDeleteinvalidemail (string) (string, error)

    CreateEmailListblockemail (*CreateEmailListblockemailInput) (string, error)

    CreateEmailListspamemail (*CreateEmailListspamemailInput) (string, error)

    CreateEmailListbounceemail (*CreateEmailListbounceemailInput) (string, error)

    CreateEmailDeletebouncesemail (string) (string, error)

    CreateEmailListinvalidemail (*CreateEmailListinvalidemailInput) (string, error)

    CreateEmailListunsubscribedemail (*CreateEmailListunsubscribedemailInput) (string, error)

    CreateEmailDeleteunsubscribedemail (string) (string, error)

    CreateEmailAddunsubscribesemail (string) (string, error)

    CreateEmailDeleteblocksemail (string) (string, error)

    CreateEmailDeletespamemail (string) (string, error)

    CreateEmailSendemails (*CreateEmailSendemailsInput) (string, error)
}

/*
 * Factory for the EMAIL interaface returning EMAIL_IMPL
 */
func NewEMAIL(config configuration_pkg.CONFIGURATION) *EMAIL_IMPL {
    client := new(EMAIL_IMPL)
    client.config = config
    return client
}
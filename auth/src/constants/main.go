package constants

import (
	"fmt"
	"html"
	"time"
)

const (
	AUTH_AGE                         = 15 * time.Minute
	OTP_AGE                          = time.Minute
	REFRESH_AGE                      = 7 * 24 * time.Hour
	VerificationLinkValidityMinutes  = 60
	ResetPasswordLinkValidityMinutes = VerificationLinkValidityMinutes
	MaxLoginAttempts                 = 3
	LoginLockDuration                = time.Hour
)

const (
	ForAuth = iota + 1
	ForOTP
	ForRefresh
)

const (
	PRODUCTION = "PRODUCTION"
)

const (
	AUTH_KEY = "auth_key"
)

// Email colors mirror client/src/theme/theme.ts: cream #F7F4EE, paper #FFFFFF,
// mist #DCE8FB, blue #3D6BD4, sky #8FB3F0, amber #F0A63B, ink #222A3B,
// slate #5C6B84, navy #253350.

func BuildVerificationEmailBody(username, verificationURL string) string {
	const verificationEmailTemplate = `<!DOCTYPE html>
	<html lang="en">
	<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Verify your Amarolio account</title>
	</head>
	<body style="margin:0;padding:0;background-color:#F7F4EE;font-family:'Poppins','Open Sans','Helvetica Neue',Arial,sans-serif;">
  <table role="presentation" width="100%%" cellpadding="0" cellspacing="0" border="0" style="background-color:#F7F4EE;padding:24px 12px;">
    <tr>
      <td align="center">
        <table role="presentation" width="100%%" cellpadding="0" cellspacing="0" border="0" style="max-width:520px;background-color:#FFFFFF;border-radius:22px;overflow:hidden;border:1px solid #DCE8FB;">

          <!-- Header -->
          <tr>
            <td align="center" style="background-color:#3D6BD4;background-image:linear-gradient(160deg,#8FB3F0 0%%,#3D6BD4 100%%);padding:30px 20px 24px 20px;">
              <div style="width:52px;height:52px;line-height:52px;border-radius:16px;background-color:#FFFFFF;color:#3D6BD4;font-size:26px;font-weight:bold;text-align:center;">A</div>
              <h1 style="margin:14px 0 0 0;font-size:26px;font-weight:bold;color:#FFFFFF;letter-spacing:1px;">Amarolio</h1>
              <p style="margin:6px 0 0 0;font-size:13px;color:#DCE8FB;">one account for every Amarolio service</p>
            </td>
          </tr>

          <!-- Accent stripe -->
          <tr>
            <td style="height:5px;line-height:5px;font-size:0;background-color:#F0A63B;">&nbsp;</td>
          </tr>

          <!-- Body -->
          <tr>
            <td style="padding:32px 32px 8px 32px;color:#222A3B;">
              <p style="margin:0 0 14px 0;font-size:20px;color:#3D6BD4;font-weight:bold;">Hello, %[1]s</p>
              <p style="margin:0 0 22px 0;font-size:16px;line-height:1.7;color:#5C6B84;">
                Welcome to Amarolio! Confirm your email address to activate your account and use it across every Amarolio service.
              </p>
            </td>
          </tr>

          <!-- Button -->
          <tr>
            <td align="center" style="padding:0 32px;">
              <table role="presentation" cellpadding="0" cellspacing="0" border="0">
                <tr>
                  <td align="center" bgcolor="#3D6BD4" style="background-color:#3D6BD4;border-radius:999px;">
                    <a href="%[2]s" target="_blank" style="display:inline-block;padding:15px 42px;font-family:'Poppins','Open Sans','Helvetica Neue',Arial,sans-serif;font-size:17px;font-weight:bold;letter-spacing:0.2px;color:#FFFFFF;text-decoration:none;">Verify my account</a>
                  </td>
                </tr>
              </table>
            </td>
          </tr>

          <!-- Fallback link + notes -->
          <tr>
            <td style="padding:26px 32px 30px 32px;color:#222A3B;">
              <p style="margin:0 0 8px 0;font-size:14px;line-height:1.6;color:#5C6B84;">
                Button not working? Copy and paste this link into your browser:
              </p>
              <p style="margin:0 0 18px 0;font-size:13px;line-height:1.5;word-break:break-all;">
                <a href="%[2]s" target="_blank" style="color:#3D6BD4;text-decoration:underline;">%[2]s</a>
              </p>
              <p style="margin:0 0 12px 0;font-size:15px;line-height:1.6;">
                This link will expire in <strong>%[3]d minutes</strong>.
              </p>
              <p style="margin:0;font-size:14px;line-height:1.6;color:#5C6B84;">
                If you didn't create an Amarolio account, you can safely ignore this email.
              </p>
            </td>
          </tr>

          <!-- Bottom bar -->
          <tr>
            <td align="center" style="background-color:#253350;padding:16px 20px;">
              <p style="margin:0;font-size:12px;color:#DCE8FB;">
                Sent by the Amarolio team
              </p>
            </td>
          </tr>

        </table>
      </td>
    </tr>
  </table>
	</body>
	</html>`
	return fmt.Sprintf(verificationEmailTemplate,
		html.EscapeString(username),        // %[1]s
		html.EscapeString(verificationURL), // %[2]s
		VerificationLinkValidityMinutes,    // %[3]d
	)
}

func BuildResetPasswordEmailBody(username, resetURL string) string {
	const resetPasswordEmailTemplate = `<!DOCTYPE html>
	<html lang="en">
	<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Reset your Amarolio password</title>
	</head>
	<body style="margin:0;padding:0;background-color:#F7F4EE;font-family:'Poppins','Open Sans','Helvetica Neue',Arial,sans-serif;">
  <table role="presentation" width="100%%" cellpadding="0" cellspacing="0" border="0" style="background-color:#F7F4EE;padding:24px 12px;">
    <tr>
      <td align="center">
        <table role="presentation" width="100%%" cellpadding="0" cellspacing="0" border="0" style="max-width:520px;background-color:#FFFFFF;border-radius:22px;overflow:hidden;border:1px solid #DCE8FB;">

          <!-- Header -->
          <tr>
            <td align="center" style="background-color:#3D6BD4;background-image:linear-gradient(160deg,#8FB3F0 0%%,#3D6BD4 100%%);padding:30px 20px 24px 20px;">
              <div style="width:52px;height:52px;line-height:52px;border-radius:16px;background-color:#FFFFFF;color:#3D6BD4;font-size:26px;font-weight:bold;text-align:center;">A</div>
              <h1 style="margin:14px 0 0 0;font-size:26px;font-weight:bold;color:#FFFFFF;letter-spacing:1px;">Amarolio</h1>
              <p style="margin:6px 0 0 0;font-size:13px;color:#DCE8FB;">one account for every Amarolio service</p>
            </td>
          </tr>

          <!-- Accent stripe -->
          <tr>
            <td style="height:5px;line-height:5px;font-size:0;background-color:#F0A63B;">&nbsp;</td>
          </tr>

          <!-- Body -->
          <tr>
            <td style="padding:32px 32px 8px 32px;color:#222A3B;">
              <p style="margin:0 0 14px 0;font-size:20px;color:#3D6BD4;font-weight:bold;">Hello, %[1]s</p>
              <p style="margin:0 0 22px 0;font-size:16px;line-height:1.7;color:#5C6B84;">
                We received a request to reset the password for your Amarolio account. Click the button below to choose a new one.
              </p>
            </td>
          </tr>

          <!-- Button -->
          <tr>
            <td align="center" style="padding:0 32px;">
              <table role="presentation" cellpadding="0" cellspacing="0" border="0">
                <tr>
                  <td align="center" bgcolor="#3D6BD4" style="background-color:#3D6BD4;border-radius:999px;">
                    <a href="%[2]s" target="_blank" style="display:inline-block;padding:15px 42px;font-family:'Poppins','Open Sans','Helvetica Neue',Arial,sans-serif;font-size:17px;font-weight:bold;letter-spacing:0.2px;color:#FFFFFF;text-decoration:none;">Reset my password</a>
                  </td>
                </tr>
              </table>
            </td>
          </tr>

          <!-- Fallback link + notes -->
          <tr>
            <td style="padding:26px 32px 30px 32px;color:#222A3B;">
              <p style="margin:0 0 8px 0;font-size:14px;line-height:1.6;color:#5C6B84;">
                Button not working? Copy and paste this link into your browser:
              </p>
              <p style="margin:0 0 18px 0;font-size:13px;line-height:1.5;word-break:break-all;">
                <a href="%[2]s" target="_blank" style="color:#3D6BD4;text-decoration:underline;">%[2]s</a>
              </p>
              <p style="margin:0 0 12px 0;font-size:15px;line-height:1.6;">
                This link will expire in <strong>%[3]d minutes</strong>.
              </p>
              <p style="margin:0;font-size:14px;line-height:1.6;color:#5C6B84;">
                If you didn't request a password reset, you can safely ignore this email — your password won't change.
              </p>
            </td>
          </tr>

          <!-- Bottom bar -->
          <tr>
            <td align="center" style="background-color:#253350;padding:16px 20px;">
              <p style="margin:0;font-size:12px;color:#DCE8FB;">
                Sent by the Amarolio team
              </p>
            </td>
          </tr>

        </table>
      </td>
    </tr>
  </table>
	</body>
	</html>`
	return fmt.Sprintf(resetPasswordEmailTemplate,
		html.EscapeString(username),      // %[1]s
		html.EscapeString(resetURL),      // %[2]s
		ResetPasswordLinkValidityMinutes, // %[3]d
	)
}

func BuildOTPEmailBody(username, otp string) string {
	const otpEmailTemplate = `<!DOCTYPE html>
	<html lang="en">
	<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Your Amarolio sign-in code</title>
	</head>
	<body style="margin:0;padding:0;background-color:#F7F4EE;font-family:'Poppins','Open Sans','Helvetica Neue',Arial,sans-serif;">
  <table role="presentation" width="100%%" cellpadding="0" cellspacing="0" border="0" style="background-color:#F7F4EE;padding:24px 12px;">
    <tr>
      <td align="center">
        <table role="presentation" width="100%%" cellpadding="0" cellspacing="0" border="0" style="max-width:520px;background-color:#FFFFFF;border-radius:22px;overflow:hidden;border:1px solid #DCE8FB;">

          <!-- Header -->
          <tr>
            <td align="center" style="background-color:#3D6BD4;background-image:linear-gradient(160deg,#8FB3F0 0%%,#3D6BD4 100%%);padding:30px 20px 24px 20px;">
              <div style="width:52px;height:52px;line-height:52px;border-radius:16px;background-color:#FFFFFF;color:#3D6BD4;font-size:26px;font-weight:bold;text-align:center;">A</div>
              <h1 style="margin:14px 0 0 0;font-size:26px;font-weight:bold;color:#FFFFFF;letter-spacing:1px;">Amarolio</h1>
              <p style="margin:6px 0 0 0;font-size:13px;color:#DCE8FB;">one account for every Amarolio service</p>
            </td>
          </tr>

          <!-- Accent stripe -->
          <tr>
            <td style="height:5px;line-height:5px;font-size:0;background-color:#F0A63B;">&nbsp;</td>
          </tr>

          <!-- Body -->
          <tr>
            <td style="padding:32px 32px 8px 32px;color:#222A3B;">
              <p style="margin:0 0 14px 0;font-size:20px;color:#3D6BD4;font-weight:bold;">Hello, %[1]s</p>
              <p style="margin:0 0 22px 0;font-size:16px;line-height:1.7;color:#5C6B84;">
                Use the one-time passcode below to continue signing in to your Amarolio account.
              </p>
            </td>
          </tr>

          <!-- OTP box -->
          <tr>
            <td align="center" style="padding:0 32px;">
              <table role="presentation" cellpadding="0" cellspacing="0" border="0" style="background-color:#DCE8FB;border:2px solid #3D6BD4;border-radius:14px;">
                <tr>
                  <td align="center" style="padding:18px 34px;">
                    <div style="font-size:12px;letter-spacing:3px;color:#3D6BD4;font-weight:bold;text-transform:uppercase;margin-bottom:8px;">Your passcode</div>
                    <div style="font-family:'Courier New',Courier,monospace;font-size:38px;font-weight:bold;letter-spacing:10px;color:#222A3B;">%[2]s</div>
                  </td>
                </tr>
              </table>
            </td>
          </tr>

          <!-- Footer text -->
          <tr>
            <td style="padding:26px 32px 30px 32px;color:#222A3B;">
              <p style="margin:0 0 12px 0;font-size:15px;line-height:1.6;">
                This code will expire in <strong>%[3]d minutes</strong>.
              </p>
              <p style="margin:0;font-size:14px;line-height:1.6;color:#5C6B84;">
                If you didn't request this, you can safely ignore this email.
                For your security, never share this code with anyone.
              </p>
            </td>
          </tr>

          <!-- Bottom bar -->
          <tr>
            <td align="center" style="background-color:#253350;padding:16px 20px;">
              <p style="margin:0;font-size:12px;color:#DCE8FB;">
                Sent by the Amarolio team
              </p>
            </td>
          </tr>

        </table>
      </td>
    </tr>
  </table>
	</body>
	</html>`
	return fmt.Sprintf(otpEmailTemplate,
		html.EscapeString(username), // %[1]s
		html.EscapeString(otp),      // %[2]s
		int(OTP_AGE.Minutes()),      // %[3]d
	)
}

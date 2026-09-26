package constants

import (
	"fmt"
	"html"
	"time"
)

const (
	AUTH_AGE                        = 15 * time.Minute
	OTP_AGE                         = time.Minute
	REFRESH_AGE                     = 7 * 24 * time.Hour
	VerificationLinkValidityMinutes = 60
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

func BuildVerificationEmailBody(username, verificationURL string) string {
	const verificationEmailTemplate = `<!DOCTYPE html>
	<html lang="en">
	<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Verify your Amary account</title>
	</head>
	<body style="margin:0;padding:0;background-color:#F3E2C7;font-family:Georgia,'Times New Roman',serif;">
  <table role="presentation" width="100%%" cellpadding="0" cellspacing="0" border="0" style="background-color:#F3E2C7;padding:24px 12px;">
    <tr>
      <td align="center">
        <table role="presentation" width="100%%" cellpadding="0" cellspacing="0" border="0" style="max-width:520px;background-color:#FFF6E9;border-radius:14px;overflow:hidden;border:1px solid #E6CFA8;">

          <!-- Header -->
          <tr>
            <td align="center" style="background-color:#C2571A;padding:28px 20px 22px 20px;">
              <div style="font-size:26px;letter-spacing:8px;line-height:1;">🍂 🍁 🍂</div>
              <h1 style="margin:12px 0 0 0;font-size:32px;font-weight:bold;color:#FFF6E9;letter-spacing:2px;">amary</h1>
              <p style="margin:4px 0 0 0;font-size:13px;color:#FAD9B0;font-style:italic;">shorten your links, gather your leaves</p>
            </td>
          </tr>

          <!-- Accent stripe -->
          <tr>
            <td style="height:6px;line-height:6px;font-size:0;background-color:#E0A526;border-bottom:3px solid #9B2C1F;">&nbsp;</td>
          </tr>

          <!-- Body -->
          <tr>
            <td style="padding:32px 32px 8px 32px;color:#4A2C17;">
              <p style="margin:0 0 14px 0;font-size:20px;color:#9B2C1F;font-weight:bold;">Hello, %[1]s 🍁</p>
              <p style="margin:0 0 22px 0;font-size:16px;line-height:1.6;">
                Welcome to Amary! The leaves are turning and it's time to confirm your email address.
                Tap the button below to verify your account.
              </p>
            </td>
          </tr>

          <!-- Button -->
          <tr>
            <td align="center" style="padding:0 32px;">
              <table role="presentation" cellpadding="0" cellspacing="0" border="0">
                <tr>
                  <td align="center" bgcolor="#C2571A" style="background-color:#C2571A;border-radius:10px;border-bottom:4px solid #9B2C1F;">
                    <a href="%[2]s" target="_blank" style="display:inline-block;padding:15px 38px;font-family:Georgia,'Times New Roman',serif;font-size:18px;font-weight:bold;letter-spacing:1px;color:#FFF6E9;text-decoration:none;">🍁 Verify my account</a>
                  </td>
                </tr>
              </table>
            </td>
          </tr>

          <!-- Fallback link + notes -->
          <tr>
            <td style="padding:26px 32px 30px 32px;color:#4A2C17;">
              <p style="margin:0 0 8px 0;font-size:14px;line-height:1.6;color:#7A5A40;">
                Button not working? Copy and paste this link into your browser:
              </p>
              <p style="margin:0 0 18px 0;font-size:13px;line-height:1.5;word-break:break-all;">
                <a href="%[2]s" target="_blank" style="color:#9B2C1F;text-decoration:underline;">%[2]s</a>
              </p>
              <p style="margin:0 0 12px 0;font-size:15px;line-height:1.6;">
                🕰️ This link will fall away in <strong>%[3]d minutes</strong>.
              </p>
              <p style="margin:0;font-size:14px;line-height:1.6;color:#7A5A40;">
                If you didn't create an Amary account, you can safely ignore this email.
              </p>
            </td>
          </tr>

          <!-- Bottom bar -->
          <tr>
            <td align="center" style="background-color:#4A2C17;padding:16px 20px;">
              <p style="margin:0;font-size:12px;color:#E6CFA8;">
                🍂 Sent with warmth by the Amary team 🍂
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

func BuildOTPEmailBody(username, otp string) string {
	const otpEmailTemplate = `<!DOCTYPE html>
	<html lang="en">
	<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Your Amary verification code</title>
	</head>
	<body style="margin:0;padding:0;background-color:#F3E2C7;font-family:Georgia,'Times New Roman',serif;">
  <table role="presentation" width="100%%" cellpadding="0" cellspacing="0" border="0" style="background-color:#F3E2C7;padding:24px 12px;">
    <tr>
      <td align="center">
        <table role="presentation" width="100%%" cellpadding="0" cellspacing="0" border="0" style="max-width:520px;background-color:#FFF6E9;border-radius:14px;overflow:hidden;border:1px solid #E6CFA8;">

          <!-- Header -->
          <tr>
            <td align="center" style="background-color:#C2571A;padding:28px 20px 22px 20px;">
              <div style="font-size:26px;letter-spacing:8px;line-height:1;">🍂 🍁 🍂</div>
              <h1 style="margin:12px 0 0 0;font-size:32px;font-weight:bold;color:#FFF6E9;letter-spacing:2px;">amary</h1>
              <p style="margin:4px 0 0 0;font-size:13px;color:#FAD9B0;font-style:italic;">shorten your links, gather your leaves</p>
            </td>
          </tr>

          <!-- Accent stripe -->
          <tr>
            <td style="height:6px;line-height:6px;font-size:0;background-color:#E0A526;border-bottom:3px solid #9B2C1F;">&nbsp;</td>
          </tr>

          <!-- Body -->
          <tr>
            <td style="padding:32px 32px 8px 32px;color:#4A2C17;">
              <p style="margin:0 0 14px 0;font-size:20px;color:#9B2C1F;font-weight:bold;">Hello, %s 🍁</p>
              <p style="margin:0 0 22px 0;font-size:16px;line-height:1.6;">
                The leaves are turning and it's time to verify your account.
                Use the one-time passcode below to continue with Amary.
              </p>
            </td>
          </tr>

          <!-- OTP box -->
          <tr>
            <td align="center" style="padding:0 32px;">
              <table role="presentation" cellpadding="0" cellspacing="0" border="0" style="background-color:#FDEBD0;border:2px dashed #C2571A;border-radius:12px;">
                <tr>
                  <td align="center" style="padding:18px 34px;">
                    <div style="font-size:12px;letter-spacing:3px;color:#9B2C1F;text-transform:uppercase;margin-bottom:8px;">Your passcode</div>
                    <div style="font-family:'Courier New',Courier,monospace;font-size:38px;font-weight:bold;letter-spacing:10px;color:#4A2C17;">%s</div>
                  </td>
                </tr>
              </table>
            </td>
          </tr>

          <!-- Footer text -->
          <tr>
            <td style="padding:26px 32px 30px 32px;color:#4A2C17;">
              <p style="margin:0 0 12px 0;font-size:15px;line-height:1.6;">
                🕰️ This code will fall away in <strong>%d minutes</strong>.
              </p>
              <p style="margin:0;font-size:14px;line-height:1.6;color:#7A5A40;">
                If you didn't request this, you can safely ignore this email.
                For your security, never share this code with anyone.
              </p>
            </td>
          </tr>

          <!-- Bottom bar -->
          <tr>
            <td align="center" style="background-color:#4A2C17;padding:16px 20px;">
              <p style="margin:0;font-size:12px;color:#E6CFA8;">
                🍂 Sent with warmth by the Amary team 🍂
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
		html.EscapeString(username),
		html.EscapeString(otp),
		int(OTP_AGE),
	)
}

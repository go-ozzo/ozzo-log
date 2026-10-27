# Security Policy

## Supported Versions

No version has been tagged yet; the first tagged release will be v1.0.0. Until then, and afterwards for the latest
release, fixes land on `master`.

Only the module path `github.com/go-ozzo/ozzo-log` receives fixes.

## Reporting a Vulnerability

**Do not open a public GitHub issue for security vulnerabilities.**

Report it privately with **Report a vulnerability** on the repository's
[Security tab](https://github.com/go-ozzo/ozzo-log/security), which opens a private advisory draft:
https://github.com/go-ozzo/ozzo-log/security/advisories/new

Please include:

- a description of the vulnerability and its impact;
- a minimal reproduction (logger and target setup, and the call that triggers it);
- the affected versions or commits you checked.

### Response Timeline

- **Initial response**: within 72 hours
- **Fix and disclosure**: coordinated with the reporter; credit is given unless you prefer otherwise

Published advisories: https://github.com/go-ozzo/ozzo-log/security/advisories

## Security Considerations

ozzo-log writes what the application hands it to files, the network and mail. Things to keep in mind:

1. **Log content** — messages are written as they are, including newlines and control characters. Do not log
   secrets (passwords, tokens, keys), and do not put untrusted input into a message without escaping it: a value with
   a line break can forge extra log lines.
2. **FileTarget** — `FileName` is used exactly as configured; the library does not check it against a base
   directory, so take it from trusted configuration, never from user input. Log files are created with mode `0660`
   (before the process umask).
3. **MailTarget** — `Username` and `Password` are used for SMTP `PLAIN` authentication, which Go's `net/smtp` sends
   only over TLS or to localhost. Keep the password out of the source code, and keep sensitive data out of the
   messages: they are sent as mail.
4. **NetworkTarget** — messages travel over the configured connection as plain text, without TLS. Use it on a
   trusted network or through a tunnel.
5. **ErrorWriter** — errors of the targets (a failed write, a failed mail) go to `Logger.ErrorWriter`. Point it at a
   place the operator reads.

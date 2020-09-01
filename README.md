# Pwned Search - Go Implementation

Checks the [HaveIBeenPwned API](https://haveibeenpwned.com/API/v2) to see if
your password has been compromised, using the [range endpoint](https://haveibeenpwned.com/API/v2#SearchingPwnedPasswordsByRange).
Inspired by [https://github.com/mikepound/pwned-search](https://github.com/mikepound/pwned-search).

# Usage

Clone it, then:
```
docker build -t pwned
docker run --rm pwned badpassword
```

# License

MIT

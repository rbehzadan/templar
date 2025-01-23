# Custom Functions in Templar

Templar includes several **custom** template functions in addition to the [Sprig functions](https://github.com/Masterminds/sprig). Below are the functions you can use in your templates:

---

## `title`

```go
title "some string"
```

- Converts the input string to title case (each word capitalized).
- Internally uses `golang.org/x/text/cases` with `language.AmericanEnglish`.

**Example**:
```bash
echo "Title: {{title \"hello world\"}}" | ./templar
# Output: Title: Hello World
```

---

## `split`

```go
split SEPARATOR STRING
```

- Splits the given `STRING` around each instance of `SEPARATOR`.
- Returns a slice of substrings.

**Example**:
```bash
echo "Splits: {{split \",\" \"a,b,c\"}}" | ./templar
# Output: Splits: [a b c]
```

---

## `domain`

```go
domain
```

- Extracts the domain portion from your machine’s hostname.
- If the hostname is something like `myhost.example.com`, it returns `example.com`.
- If the hostname or parsing fails, returns the default string `"exmaple.ir"` (typo preserved from code).

**Example**:
```bash
echo "Domain: {{domain}}" | ./templar
# Output: Domain: example.com
```

---

## `genpw`

```go
genpw LENGTH
```

- Generates a random password of length `LENGTH` using a predefined character set:
  - `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()`
- If password generation fails (very unlikely), it returns a fallback password.

**Example**:
```bash
echo "Random Password: {{genpw 16}}" | ./templar
# Output: Random Password: 5iw&fT!OHQM3C0Jk
```

---

## Notes

- All these functions can be combined with Sprig functions.  
- For a full list of functions from Sprig, see the [Sprig documentation](https://masterminds.github.io/sprig/).


# Adding hand-written code later

Nothing here is hand-written yet: every SDK file is generated. This document records how hand-written code will be added when it is needed, using only mainline openapi-generator features ([customization docs](https://openapi-generator.tech/docs/customization/)), so that regeneration keeps working and never silently discards it.

## The three mechanisms

1. **`.openapi-generator-ignore`** (gitignore syntax, one per language directory). A file whose path matches is never written or overwritten by the generator. Use it to protect hand-written files.
2. **`templateDir`** (config key, or `-t` on the command line). A directory of mustache templates that override the generator's built-in templates of the same name.
3. **`files`** (config node). Adds or replaces supporting files from custom templates: each entry names a template, a destination folder and a destination file name.

## Preferred order

1. **Composition: hand-written files in their own paths.** Put the code where the generator never writes, for example `ruby/lib/pharmaobservatory/ext/`, `go/ext/` or `rust/src/ext/`, and list the paths in that language's `.openapi-generator-ignore`. Anything that wraps or extends the generated client (token refresh, retry policy, convenience methods, pagination helpers) belongs here. Wire it in with the smallest hook the language allows, for example a `require` from a protected file, and, if the generated entry point (`lib/pharmaobservatory.rb`, `lib.rs`) has to mention it, use step 2 for that one line.
2. **Small template overrides** under `templates/<language>/`, only for a change that must live inside a generated file (an extra dependency in the gemspec or `Cargo.toml`, a new line in the entry point). Point `templateDir` at that directory in the language's `openapi-config.json`. Copy the template from the pinned generator version, change the minimum, and say in the template's first comment which upstream template and version it came from, because a generator upgrade must re-diff it.
3. **Never edit a generated file in place.** The next regeneration overwrites it.

## Rules that keep regeneration safe

- Every hand-written path is listed in `.openapi-generator-ignore`. `bin/generate` deletes the files the generator wrote last time (from `.openapi-generator/FILES`) before regenerating, and skips the entries it finds listed literally in the ignore file, so list protected files by exact path (a glob protects the file from being written but does not stop the clean-up step from removing it: protect by exact path, or keep the code in a directory that is not in `FILES`).
- Hand-written tests live beside the hand-written code, under paths the generator does not own, and run in the existing CI job for that language.
- Today's post-generation steps in `bin/generate` (the gemspec metadata and file list, the Cargo metadata and TLS feature) exist because the generator has no option for them. A template override is the cleaner home for each once template overrides are in use; move them only when there is a reason to touch them.
- Bump the generator version in its own pull request and regenerate in that PR, so a template override that no longer matches shows up as a diff.

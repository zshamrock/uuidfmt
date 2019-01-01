# uuidfmt

[![Build Status](https://travis-ci.org/zshamrock/uuidfmt.svg?branch=master)](https://travis-ci.org/zshamrock/uuidfmt)
[![Snap Status](https://build.snapcraft.io/badge/zshamrock/uuidfmt.svg)](https://build.snapcraft.io/user/zshamrock/uuidfmt)

Really simple CLI tool to format UUID into the canonical presentation.

## Installation

Use the `go` command:

	$ go get github.com/zshamrock/uuidfmt

Or using `snap`:

    $ snap install uuidfmt
    
[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-black.svg)](https://snapcraft.io/uuidfmt)

## Usage

    $ uuidfmt 3A2DD5E0D2C04F13A3E2F600C9530793

![uuidfmt](uuidfmt.gif)

## Options

    -r/--reverse - apply the reverse formatting

i.e. `3a2dd5e0-d2c0-4f13-a3e2-f600c9530793` `->` `3A2DD5E0D2C04F13A3E2F600C9530793`.

## Copyright

Copyright (C) 2018 by Aliaksandr Kazlou.

uuidfmt is released under MIT License.
See [LICENSE](https://github.com/zshamrock/uuidfmt/blob/master/LICENSE) for details.

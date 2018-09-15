```
     _   _ ____  __  __ _  ___ _
    | | | / ___||  \/  | |/ (_) |_
    | |_| \___ \| |\/| | ' /| | __|
    |  _  |___) | |  | | . \| | |_
    |_| |_|____/|_|  |_|_|\_\_|\__|
```

[![Maintenance](https://img.shields.io/maintenance/yes/2018.svg)]() [![License](https://img.shields.io/github/license/hashlab/hsmkit.svg)](https://github.com/hashlab/hsmkit/blob/master/LICENSE) [![GoDoc](https://godoc.org/github.com/hashlab/hsmkit?status.svg)](https://godoc.org/github.com/hashlab/hsmkit/lib) [![Twitter Follow](https://img.shields.io/twitter/follow/chrisenytc.svg?style=social&label=Follow)](http://twitter.com/chrisenytc) [![Twitter URL](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Awesome%20https://github.com/hashlab/hsmkit%20via%20@chrisenytc)

> A golang tool to generate keys required for using an HSM

## Installation

```bash
$ go get -u github.com/hashlab/hsmkit
```

## Usage

Generating a component key

```bash
$ hsmkit generate-component --force-odd

Component Key: 1CAD7AD361548AA1628FAB31A71C6B80
KCV: 46F084
```

Combining 3 component keys to generate a combined key

```bash
$ hsmkit combine

? Enter the component key 1: ********************************
? Enter the component key 2: ********************************
? Enter the component key 3: ********************************

Combined Key: DCCB3B10D6402902F15B1A151CF16723
Final KCV: 315381
```

## Getting Started

1º Clone hsmkit repo

```bash
$ git clone git@github.com:hashlab/hsmkit.git
```

2º Enter in hsmkit directory
```bash
$ cd hsmkit
```

3º Run the builder
```bash
$ make build
```

4º Run the CLI
```bash
$ ./bin/cli
```

## Conventions of commit messages

Addding files on repo

```bash
git commit -m "Add filename"
```

Updating files on repo

```bash
git commit -m "Update filename, filename2, filename3"
```

Removing files on repo

```bash
git commit -m "Remove filename"
```

Renaming files on repo

```bash
git commit -m "Rename filename"
```

Fixing errors and issues on repo

```bash
git commit -m "Fixed #issuenumber Message about this fix"
```

Adding features on repo

```bash
git commit -m "Add feature: nameoffeature Message about this feature"
```

Updating features on repo

```bash
git commit -m "Update feature: nameoffeature Message about this update"
```

Removing features on repo

```bash
git commit -m "Remove feature: nameoffeature Message about this"
```

## Contributing

Bug reports and pull requests are welcome on GitHub at [https://github.com/hashlab/hsmkit](https://github.com/hashlab/hsmkit). This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

1. Fork it [hashlab/hsmkit](https://github.com/hashlab/hsmkit/fork)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am "Add some feature"`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## Support
If you have any problem or suggestion please open an issue [here](https://github.com/hashlab/hsmkit/issues).

## License 

Check [here](LICENSE).

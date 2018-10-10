```
     ____  _   _ _  ______ _____ _  ___ _
    |  _ \| | | | |/ /  _ \_   _| |/ (_) |_
    | | | | | | | ' /| |_) || | | ' /| | __|
    | |_| | |_| | . \|  __/ | | | . \| | |_
    |____/ \___/|_|\_\_|    |_| |_|\_\_|\__|
```

[![Maintenance](https://img.shields.io/maintenance/yes/2018.svg)]() [![License](https://img.shields.io/github/license/hashlab/dukptkit.svg)](https://github.com/hashlab/dukptkit/blob/master/LICENSE) [![GoDoc](https://godoc.org/github.com/hashlab/dukptkit?status.svg)](https://godoc.org/github.com/hashlab/dukptkit/lib) [![Twitter Follow](https://img.shields.io/twitter/follow/chrisenytc.svg?style=social&label=Follow)](http://twitter.com/chrisenytc) [![Twitter URL](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Awesome%20https://github.com/hashlab/dukptkit%20via%20@chrisenytc)

> A golang tool to generate keys required for using 3DES-DUKPT

## Installation

```bash
$ go get -u github.com/hashlab/dukptkit
```

## Usage

Generating a BDK identifier (KSI)

```bash
$ dukptkit generate-ksi

BDK identifier (KSI): D56D 0AFA 76
```

Generating a component key

```bash
$ dukptkit generate-component --force-odd

Component Key: AD31 4062 6EEF BCE0 E313 5B45 73EA F794
KCV: D1FD6C
Odd Parity: true
```

Combining 3 component keys to generate a combined key

```bash
$ dukptkit combine --show-combined-key

? Enter the component key 1: ********************************
? Enter the component key 2: ********************************
? Enter the component key 3: ********************************

KCV from Component Key 1: 461F68
KCV from Component Key 2: 2EC7BC
KCV from Component Key 3: B69F03

Combined Key: B329 9B2A 38BC 9D8C 3EEC 0BCE 9237 F17F
Final KCV: 9F1DBF
Odd Parity: true
```

## Getting Started

1º Clone dukptkit repo

```bash
$ git clone git@github.com:hashlab/dukptkit.git
```

2º Enter in dukptkit directory
```bash
$ cd dukptkit
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

Bug reports and pull requests are welcome on GitHub at [https://github.com/hashlab/dukptkit](https://github.com/hashlab/dukptkit). This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

1. Fork it [hashlab/dukptkit](https://github.com/hashlab/dukptkit/fork)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am "Add some feature"`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## Support
If you have any problem or suggestion please open an issue [here](https://github.com/hashlab/dukptkit/issues).

## License

Check [here](LICENSE).

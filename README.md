# X.509 Certificate Generator

This Go program generates a self-signed X.509 certificate, private key, and PFX file for testing and development purposes. 
This can be a handy utility for creating certificates for testing HTTPS, TLS, code signing, or authentication flows.

## Features

- Generates a 2048-bit RSA private key 
- Creates a self-signed certificate using the private key
- Encodes private key and certificate in PEM format
- Writes private key, certificate, and PFX file to disk 
- Outputs file names for easy testing

## Usage

Simply run the program:

```
go run main.go
```

It will generate unique file names each time using UUIDs. 

Sample output:

```
Written A1 certificate files:
Private Key: c918a919-4157-4832-8ac2-9cadffi3872b.key 
Certificate: 15db6f4a-f94e-4426-9e76-33735421211a.crt
PFX File: a87362ba-7586-4971-8f29-754834f97d9f.pfx
```

The certificate is valid for 180 days and issued to "Acme Co".

The PFX file is protected with the password "password".

## Customization

You can easily tweak the certificate details like subject, extensions, and validity period by modifying the `main.go` file.

## Backstory

I created this tool because I needed a simple way to generate self-signed certificates for a project at work.
It's amazing how much mileage you can get from the standard library crypto and x509 packages!

> This project scratches my own itch. But I think other developers and testers could find it useful too.

In the future I may add:

- More customization
- A UI or TUI for interactive use
- Support for other key types like ECC
- Additional certificate extensions like Subject Alternative Name
- Configuration through a config file or command line flags

> If you have any other feature ideas, please open an issue!

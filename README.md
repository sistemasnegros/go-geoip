<p align="center">
  <a href="" rel="noopener">
 <img height=200px src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT0ovDFqKcen5cR69OI55mrbZh3FFNBqTUrAgNDQCdO5TsNy10vdKhV6gobiakmLDezheQ&usqp=CAU" alt="Project logo"></a>
</p>

<h3 align="center">Geoip Lookup</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/kylelobo/The-Documentation-Compendium.svg)](https://github.com/kylelobo/The-Documentation-Compendium/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center"> This project write in goolang for be cli tool for geoip lookup.
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Deployment](#deployment)
- [Usage](#usage)
- [Built Using](#built_using)
- [Authors](#authors)


## 🧐 About <a name = "about"></a>

This project is used to lookup ip get source:

- Country
- City (pending)
- ASN (pending)



## 🏁 Getting Started <a name = "getting_started"></a>

### Prerequisites

```
GeoLite2-Country.mmdb
```
Download from [dev.maxmind.com](https://dev.maxmind.com/geoip/geolocate-an-ip/databases?lang=en) GeoLite2-Country.mmdb and save in /home/user/GeoLite2-Country.mmdb or in same directory with the binary


Site alternative download

```
curl -L "https://git.io/GeoLite2-Country.mmdb" -o ~/GeoLite2-Country.mmdb
```
### Installing from binary
```
curl -L "https://github.com/sistemasnegros/go-geoip/releases/download/latest/geoip" -o /usr/local/bin/geoip
```

### Installing from source

Prerequisites
```
go 1.20
```

```
git clone https://github.com/sistemasnegros/go-geoip
cd go-geoip/src
```

Once you've cloned the project, install dependencies with

```
go mod download
```


## 🎈 Usage <a name="usage"></a>

### Run from source

```bash
go run . 8.8.8.8
```

## 🚀 Deployment <a name = "deployment"></a>

To create a production version of your app:

```bash
go build -o geoip
chmod +x  geoip
```

### Run from binary
```bash
./geoip 8.8.8.8
```
output
```bash
United States => 8.8.8.8 
```

## ⛏️ Built Using <a name = "built_using"></a>

- [Go](https://go.dev/) - Programming Language.
- [GeoipLib](https://github.com/oschwald/geoip2-golang) - This library reads MaxMind GeoLite2 and GeoIP2 databases.
- [Maxmind](https://dev.maxmind.com/geoip/geolite2-free-geolocation-data?lang=en) - Geolocating an IP address using GeoIP2 and GeoLite2 databases consists of configuring a database reader and querying the database. 


## ✍️ Authors <a name = "authors"></a>

- [@sistemasnegros](https://github.com/sistemasnegros) - Idea & Initial work



# Indexer

CLI tool to index emails in text format (see example below) in the search engine [ZincSearch](https://zincsearch-docs.zinc.dev).


## Email Format

An txt email has the following structure:

```txt
Message-ID: ...
Date: ...
From: ...
To: ...
Subject: ...
Mime-Version: ...
Content-Type: ...
Content-Transfer-Encoding: ...
X-From: ...
X-To: ...
X-cc: ...
X-bcc: ...
X-Folder: ...
X-Origin: ...
X-FileName: ...
...
...
...
```

It is assumed that the entire content of the email after the "X-FileName" field is the body of the email. Examples of emails can be found in the [samples](https://github.com/timetravel-1010/indexer/tree/main/samples) folder.


### Data Structure

The structure and corresponding json fields used to parse an email are shown below. An email can contain at most, but not necessarily, the fields described below:

```go
type Email struct {
	MessageID               string          `json:"Message-Id"`
	Date                    string          `json:"Date"`
	From                    string          `json:"From"`
	To                      []*mail.Address `json:"To"`
	CC                      []*mail.Address `json:"Cc"`
	BCC                     []*mail.Address `json:"Bcc"`
	Subject                 string          `json:"Subject"`
	MimeVersion             string          `json:"Mime-Version"`
	ContentType             string          `json:"Content-Type"`
	ContentTransferEncoding string          `json:"Content-Transfer-Encoding"`
	XFrom                   string          `json:"X-From"`
	XTo                     []*mail.Address `json:"X-To"`
	Xcc                     []*mail.Address `json:"X-Cc"`
	Xbcc                    []*mail.Address `json:"X-Bcc"`
	XFolder                 string          `json:"X-Folder"`
	XOrigin                 string          `json:"X-Origin"`
	XFileName               string          `json:"X-Filename"`
	Body                    string          `json:"Body"`
}

```

## Setting up ZincSearch

First create a data folder to store the zincsearch data 

```bash
mkdir data 
```

Now grant the corresponding permissions to the data folder

```bash
chmod -R a+rwx data
```

You can use the command to run ZincSearch in a docker container using the [official image](https://zincsearch-docs.zinc.dev/installation/).

```bash
sudo docker run -v ./data:/data \
-e ZINC_DATA_PATH="/data" \
-p 4080:4080 \
-e ZINC_FIRST_ADMIN_USER=admin \
-e ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123 \
--name zincsearch public.ecr.aws/zinclabs/zincsearch:latest
```

That's it! Now you should have ZincSearch running on the port 4080, check in http://localhost:4080.

## Build from Source

Run the following command 

```
go build -o indexer .
```

You can replace `indexer` which your preferred name for the binary.

## Usage

To start indexing run the following command specifying the maildir directory which contains the text emails. 

```
./indexer  -dir <folder name>
```

This way uses the default flag values, as follows:

``` 
-dir string
	path to email directory (default "enron_mail_20110402")

-index string
	index name (default "enron")

-password string
	zincsearch password (default "Complexpass#123")

-port string
	zincsearch host port (default "4080")

-user string
	zincsearch username (default "admin")

-zincurl string
	zincsearch host url (default "localhost")
```

So, the default directory is `enron_mail_20110402` which corresponds to [this](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz) email database.

Other available flags are: 

```
-custom
	use custom implementation instead of std (net/mail) library

-type string
	request payload type (default "_doc")
```

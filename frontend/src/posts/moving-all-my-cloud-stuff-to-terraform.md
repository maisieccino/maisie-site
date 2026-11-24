---
title: Moving my cloud to Terraform
authors:
  - Maisie Bell
publishedTime: 2020-04-25T13:30:00.000+00:00
modifiedTime: 2020-10-04T12:40:31.000+00:00
featureImage: /src/assets/img/posts/imported/Untitled_Artwork-2.png
featureImageAlt: ""
tags: []
excerpt: I want it all. The terrifying lows, the dizzying highs. Here's what
  happened when I moved my infrastructure to Terraform.
readingTime: 4

---

I want it all. The terrifying lows, the dizzying highs.

A couple of weeks ago, I had an idea: why not make all of my cloud resources be defined in Terraform code? For those not in the know, [Terraform](https://www.terraform.io/?ref=blog.mbell.dev) is a technology that lets you define all of your infrastructure in code, made by a company called Hashicorp.

For the most part it's surprisingly intuitive! The language used, called "HCL", has a syntax not too unfamiliar from Go and C. It lets you define loops, local variables, and create "modules", which are reusable blocks of infrastructure. It also comes with a nice templating engine for generating files.

One of the first things to think about when using Terraform is where to store your state. Terraform has a "state" concept, which is your project's current snapshot of what it thinks your infrastructure estate should look like. This is updated whenever you deploy, or "apply" your code. For most individuals, this is just stored on your computer\*, but I didn't think that was good enough. You could store the state in a cloud bucket (like Amazon S3 or Google Cloud Storage), but that would also involve extra set up.

Fortunately, Hashicorp offer a hosted version of Terraform which is free for invididuals. As well as looking after your Terraform state, it can perform Terraform runs remotely, meaning that you can just push code to Git and they will generate the "plan" (a bit like a dry run, updates the state with what the infrastructure is like right now, and calculates what changes will be made), and with one click, actually applies the plan too.

![](/src/assets/img/posts/imported/image--1-.png)

The Terraform development lifecycle. The "backend" is where the Terraform state is stored, be this a local file, a cloud storage bucket, or a hosted service.

## Building new infrastructure

The first thing I did was to jump in and start building some new infrastructure in Terraform. I had joked to a friend that we should play Minecraft again as a way of coping with the quarantine, and to my surprise they were up for it. Perfect excuse to flex my Terraform.

I needed a server, storage device, an SSL certificate, and a DNS record pointing at this new server. So I got to work. The results are in [GitHub](https://github.com/mbellgb/infra/tree/master/modules/minecraft?ref=blog.mbell.dev).

Creating the server and load balancer was surprisingly easy work! And once I had got the hang of writing [Cloud-Init](https://cloudinit.readthedocs.io/en/latest/topics/examples.html?ref=blog.mbell.dev) files for setting up the server, things got a lot smoother.

> Cloud-init is a file format that allows you to define packages, users, and commands to run when a cloud instance is created. It can be a YAML file defining a bunch of modules, or it can be a simple shell script that is run when the instance is started for the first time.

The trickier part was sorting out my DNS records. At the start of this project, they were living in Cloudflare, however I needed the top level domain to be managed by DigitalOcean for them to automatically provision SSL certificates for me. This wasn't too bad, I could define each DNS record as new `digitalocean_record` resources and then when I was ready, I updated the nameserver to point at DigitalOcean's nameservers.

Then I decided to import my other DNS zones into terraform, for `mbell.me` and `mbell.dev`. I decided to keep these stored in Cloudflare, but managed in Terraform. The import process is a little bit painful if you don't manage your Terraform state manually, as the import process has to run on your local machine, and then the results sent back to the remote server. Terraform Cloud won't let you retreive sensitive project variables, so you have to redefine them again locally.

However, that wasn't the end of my problems. There seems to be a bug in Terraform that won't let you override variables locally if you have a remote backend, meaning that I had to temporarily hard-code my Cloudflare API key into the local copy of my files. Then, to make matters worse, the import process requires the 'ID' of each DNS record cloudflare has. This has to be fetched using the Cloudflare API as it is not present on the web interface. So I had to write a little script to fetch them and run each `terraform import` command:

```shell
#!/bin/bash -e
email_header="X-Auth-Email: ${CF_EMAIL}"
auth_header="X-Auth-Key: ${CF_TOKEN}"
zone_id="redacted"
curl \
    -H "${auth_header}" \
    -H "${email_header}" \
    "https://api.cloudflare.com/client/v4/zones/${zone_id}/dns_records" \
    | jq '.result | map({id: .id, name: .name, type: .type})' | jq -c ".[]" >/tmp/zones
while read -r zone; do
    id=$(jq -r .id <<< "$zone")
    name=$(jq -r .name <<< "$zone")
    type=$(jq -r .type <<< "$zone")
    echo "Terraform resource for $name ($type):"
    read resource </dev/tty
    terraform import "${resource}" "${zone_id}/${id}"
```

This script fetches the list of DNS records for the given name, and then for each one, prompt me to provide the name of the Terraform resource to import it to.

Finally, once this was done, the state was stored in Terraform Cloud, and I could plan & apply the code. Success! Nothing deleted.

## Importing existing servers

I still need to import a few other resources into Terraform, namely this blog that I'm writing on, and the bucket that contains audio files for the [NaCl podcast](https://nacl.bell.wtf/?ref=blog.mbell.dev). Hopefully now, these shouldn't be too difficult and I should be able to avoid accidentally deleting my files.

### While you're here...

Did you know you can also sign up for emails whenever I publish a new post? Not that I post very often, let's be honest. Maybe soon I'll dabble with a semi-regular newsletter or something, who knows. Subscribe below!

<https://blog.mbell.dev/signup/>

***

*\*Do not ever store your state file in version control. This file can easily contain credentials and secrets, and would be very dangerous if made public.*

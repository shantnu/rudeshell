terraform {
  cloud {
    organization = "shantnu"
    workspaces {
      name = "rudeshell"
    }
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }
}

provider "aws" {

  profile = "default"
  region  = "us-east-1"
}

resource "aws_instance" "app_server" {
  ami           = var.aws_ami_name
  instance_type = "t2.small"
  key_name      = "ec2"
  vpc_security_group_ids = ["${aws_security_group.rude-shell-sg-allow-http-ssh.id}"]

  tags = {
    Name = var.instance_name
  }
}


resource "aws_security_group" "rude-shell-sg-allow-http-ssh" {

  name        = "rude-shell-sg-allow-http-ssh"
  description = "Allow http, https, TLS inbound traffic"

  ingress {
    description = "TLS from VPC"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]

  }

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }


  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }


}

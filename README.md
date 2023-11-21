# Project hyperkv

Parse Hyper-V Key Value Pairs on Linux Guest

This is a simple library to parse Hyper-V Key Value Pairs on Linux Guest. 

All of kv files are located in `/var/lib/hyperv/` directory and maintained by `hv_kvp_daemon`.
`hv_kvp_daemon` is a daemon that runs on Linux guest and communicates with Hyper-V host to exchange data.

For more info, please refer to [Data Exchange: Using key-value pairs to share information between the host and guest on Hyper-V](https://learn.microsoft.com/en-us/previous-versions/windows/it-pro/windows-server-2012-R2-and-2012/dn798287(v=ws.11)#linux-guests)

## Usage

```shell
hyperkv -f /var/lib/hyperv/.kvp_pool_3 -o plain
```

It is not an official Microsoft project.
Similar projects [hyperkv](https://github.com/phistrom/hyperkv?tab=readme-ov-file#hyperkv)

## Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

## Trademarks

This project may contain trademarks or logos for projects, products, or services. Authorized use of Microsoft 
trademarks or logos is subject to and must follow 
[Microsoft's Trademark & Brand Guidelines](https://www.microsoft.com/en-us/legal/intellectualproperty/trademarks/usage/general).
Use of Microsoft trademarks or logos in modified versions of this project must not cause confusion or imply Microsoft sponsorship.
Any use of third-party trademarks or logos are subject to those third-party's policies.
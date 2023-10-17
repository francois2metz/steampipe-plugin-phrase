# Phrase plugin for Steampipe

Use SQL to query projects, translations and more from [phrase][].

- **[Get started →](docs/index.md)**
- Documentation: [Table definitions & examples](docs/tables)

## Quick start

Install the plugin with [Steampipe][]:

    steampipe plugin install francois2metz/phrase

## Development

To build the plugin and install it in your `.steampipe` directory

    make

Copy the default config file:

    cp config/phrase.spc ~/.steampipe/config/phrase.spc

## License

Apache 2

[steampipe]: https://steampipe.io
[phrase]: https://phrase.com

package internal



type Config struct{
	HostName string `env:"host_name" envDefault:"https://musicbrainz.org/ws/2/artist/f27ec8db-af05-4f36-916e-3d57f91ecf5e?&fmt=json&inc=url-rels+release-groups"`
}
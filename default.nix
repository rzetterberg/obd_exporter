{ go, buildGoPackage }:

let
  shortName = "obd_exporter";
  version   = "0.1.0";
in
  buildGoPackage rec {
    name          = "${shortName}-${version}";
    goPackagePath = "github.com/rzetterberg/${shortName}";
    src           = ./src;
    buildInputs   = [
      go
    ];
  }

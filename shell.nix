{ system ? builtins.currentSystem }:

let
  pkgs = import <nixpkgs> { inherit system; };
in
  with pkgs; stdenv.mkDerivation rec {
    name    = "obd_exporter-${version}";
    version = "0.2.0";

    buildInputs = [
      go
    ];
  }

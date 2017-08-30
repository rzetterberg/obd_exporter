{ system ? builtins.currentSystem }:

let
  pkgs = import <nixpkgs> { inherit system; };
in
  pkgs.callPackage ./default.nix {
    inherit (pkgs);
  }

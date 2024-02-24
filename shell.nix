{ pkgs ? import <nixpkgs> { } }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.gotools # godoc, ...
    pkgs.go-tools # staticcheck, ...
    pkgs.delve
    pkgs.gopls
    pkgs.gcc
    pkgs.gomodifytags
    pkgs.gore
    pkgs.gotests
    pkgs.gocode
    pkgs.govulncheck
    pkgs.libGL
    pkgs.mesa
    pkgs.xorg.libXcursor
    pkgs.xorg.libXi
    pkgs.xorg.libXinerama
    pkgs.xorg.libXrandr
    pkgs.xorg.libXxf86vm
    pkgs.pkg-config
    pkgs.libglvnd
    pkgs.libglibutil
    pkgs.gcc
    pkgs.bashInteractive
  ];
}

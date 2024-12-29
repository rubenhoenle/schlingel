{
  description = "An application to share work";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, nixpkgs, treefmt-nix }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };
      treefmtEval = treefmt-nix.lib.evalModule pkgs ./treefmt.nix;

      schlingel = pkgs.buildGoModule {
        name = "schlingel";
        version = "0.0.1";
        #vendorHash = pkgs.lib.fakeHash;
        vendorHash = "sha256-hg+S9rfKPpCDsFyXXrGDtHNytFQ3AIU/PJA6jCLvYdU=";
        preBuild = ''
          ${pkgs.templ}/bin/templ generate
        '';
        src = ./.;
      };

      containerImage = pkgs.dockerTools.buildLayeredImage {
        name = "ghcr.io/rubenhoenle/schlingel";
        tag = "unstable";
        config = {
          Entrypoint = [ "${schlingel}/bin/schlingel" ];
        };
      };
    in
    {
      formatter.${system} = treefmtEval.config.build.wrapper;
      packages.${system} = {
        default = schlingel;
        containerimage = containerImage;
      };
      devShells.${system}.default = pkgs.mkShell {
        buildInputs = with pkgs; [
          go
          gopls
          reflex
          templ
        ];
      };
    };

}

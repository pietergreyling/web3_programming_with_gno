# - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
# A FILE TO KEEP CONVENIENT SNIPPETS

# - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
# Installing Gno
# https://github.com/gnolang/gno/tree/master/gnovm/cmd/gno#install
  
# I will use this method:

cd [YOUR_PROJECT_ROOT_DIRECTORY]
git clone https://github.com/gnolang/gno.git
cd ./gno
make install_gno

# But one can also install gno like this:

go install github.com/gnolang/gno/gnovm/cmd/gno

# - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
# Installing Gnokey
# https://github.com/gnolang/gno/tree/master/gno.land/cmd/gnokey#install-gnokey

cd [YOUR_PROJECT_ROOT_DIRECTORY]
# Assuming the following has been executed:
#   https://github.com/gnolang/gno.git
cd ./gno
make install_gnokey

# - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
# Installing gofumpt
# https://github.com/mvdan/gofumpt

go install mvdan.cc/gofumpt@latest

# - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -



############ Mac ##########

# Install Homebrew
ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

# Install python
brew install python
# Install pip
sudo easy_install pip
# Install VirtualEnv
sudo install virtualenv
# Create a virtual environment
VirtEnv=.virtenv
virtualenv "$VirtEnv"
source "$VirtEnv/bin/activate"
pip --disable-pip-version-check install --upgrade --requirement python_requirements.txt > pip-install.log


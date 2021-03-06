############ Mac ##########

# Install Homebrew
ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

# Install python
brew install python
# Install pip
sudo easy_install pip
# Install VirtualEnv
pip install virtualenv
# Create a virtual environment
VirtEnv=.virtenv
virtualenv "$VirtEnv"
source "$VirtEnv/bin/activate"
pip --disable-pip-version-check install --upgrade --requirement python_requirements.txt > pip-install.log

# Cut video from long video
brew install ffmpeg 

brew install opencv --python27 --ffmpeg

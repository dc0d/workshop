git ls-files -m . | grep -v '_test\.go$' | grep '\.go$' | xargs git checkout HEAD --

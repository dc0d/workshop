git ls-files -m . | grep -v '_test\.exs$' | grep '\.ex$' | xargs git checkout HEAD --


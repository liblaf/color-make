if has make && has color-make; then
  function makec() {
    command make "${@}" 2>&1 | color-make
  }
fi

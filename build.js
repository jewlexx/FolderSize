// Lil javascript script to build for all the main platforms
const child = require('child_process');

const getOutput = (arch, platform) => {
  if (platform === 'windows') {
    return `fs-windows-${arch}.exe`;
  }
  return `fs-${platform}-${arch}`;
};

const build = (arch, platform) => {
  const output = getOutput(arch, platform);

  child.execSync(
    `env GOOS=${platform} GOARCH=${arch} go build -o ${output} main.go`,
  );
};

// Linux
build('amd64', 'linux');
build('386', 'linux');
build('arm64', 'linux');
build('arm', 'linux');

// Windows
build('amd64', 'windows');
build('386', 'windows');

// MacOS
build('amd64', 'darwin');
build('arm64', 'darwin');

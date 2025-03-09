import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  distDir: 'build',
  output: "export",
  trailingSlash: true,
}

export default nextConfig;

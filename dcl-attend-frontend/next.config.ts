import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  output: "export",
  distDir: "out",
  // Required for GitHub Pages deployment - turned off when deploying to cv.dean.lol
  // basePath: process.env.NODE_ENV === "production" ? "/deanlol-cv" : "/",
  images: {
    unoptimized: true,
  },
};

export default nextConfig;

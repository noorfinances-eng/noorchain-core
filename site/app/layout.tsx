import "./globals.css";
import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "NOORCHAIN — Official Site",
  description: "Social Signal Blockchain powered by PoSS."
};

export default function RootLayout({
  children
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}

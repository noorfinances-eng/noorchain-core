import "./globals.css";
import type { Metadata } from "next";
import NavBar from "../components/NavBar";
import Footer from "../components/Footer";
import { Inter } from "next/font/google";

const inter = Inter({
  subsets: ["latin"],
  display: "swap",
});

export const metadata: Metadata = {
  title: "NOORCHAIN — Social Signal Blockchain",
  description:
    "NOORCHAIN is a Social Signal Blockchain powered by PoSS. Transparent, fixed-supply, non-financial, mission-driven.",
  openGraph: {
    title: "NOORCHAIN — Social Signal Blockchain",
    description:
      "A blockchain focused on transparent social participation and curator validation. Fixed supply. No financial promises.",
    url: "https://noorchain-core-lrhx.vercel.app/",
    siteName: "NOORCHAIN",
    type: "website",
    locale: "en_US",
  },
  keywords: [
    "NOORCHAIN",
    "PoSS",
    "Social Signal Blockchain",
    "Curators",
    "Fixed Supply",
    "Transparent Governance",
  ],
  robots: "index, follow",
  icons: {
    icon: "/favicon.svg",
  },
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body
        className={`${inter.className} flex flex-col min-h-screen bg-white text-gray-900`}
      >
        <NavBar />
        <main className="flex-1">{children}</main>
        <Footer />
      </body>
    </html>
  );
}

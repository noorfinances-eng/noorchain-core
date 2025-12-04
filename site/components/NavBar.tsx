"use client";

import Link from "next/link";

export default function NavBar() {
  return (
    <nav className="w-full flex items-center justify-between px-6 py-4 border-b">
      <div className="text-xl font-bold">NOORCHAIN</div>

      <div className="flex gap-6 text-sm">
        <Link href="/">Home</Link>
        <Link href="/technology">Technology</Link>
        <Link href="/poss">PoSS</Link>
        <Link href="/curators">Curators</Link>
        <Link href="/genesis">Genesis</Link>
        <Link href="/roadmap">Roadmap</Link>
        <Link href="/docs">Docs</Link>
        <Link href="/legal">Legal</Link>
      </div>
    </nav>
  );
}

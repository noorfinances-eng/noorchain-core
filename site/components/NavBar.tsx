"use client";

import Link from "next/link";
import Image from "next/image";

export default function NavBar() {
  return (
    <nav className="w-full border-b bg-white">
      <div className="container flex items-center justify-between py-4">

        {/* Left: Logo + Name */}
        <Link href="/" className="flex items-center gap-3">
          <Image
            src="/logo.svg"
            alt="NOORCHAIN Logo"
            width={36}
            height={36}
            priority={true}
          />
          <span className="text-xl font-bold tracking-tight">NOORCHAIN</span>
        </Link>

        {/* Right: Menu */}
        <div className="flex gap-5 text-sm text-gray-700">
          <Link href="/" className="hover:text-black">Home</Link>
          <Link href="/technology" className="hover:text-black">Technology</Link>
          <Link href="/poss" className="hover:text-black">PoSS</Link>
          <Link href="/curators" className="hover:text-black">Curators</Link>
          <Link href="/genesis" className="hover:text-black">Genesis</Link>
          <Link href="/roadmap" className="hover:text-black">Roadmap</Link>
          <Link href="/docs" className="hover:text-black">Docs</Link>
          <Link href="/legal" className="hover:text-black">Legal</Link>
        </div>
        
      </div>
    </nav>
  );
}

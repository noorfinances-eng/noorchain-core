"use client";

import Link from "next/link";
import Image from "next/image";

export default function NavBar() {
  return (
    <nav className="w-full bg-white border-b border-gray-soft shadow-sm">
      <div className="container flex items-center justify-between py-4">

        {/* LEFT: LOGO + NAME */}
        <Link href="/" className="flex items-center gap-3">
          <Image
            src="/logo.svg"
            alt="NOORCHAIN Logo"
            width={40}
            height={40}
            priority
          />
          <span className="text-xl font-bold tracking-tight text-navy">
            NOORCHAIN
          </span>
        </Link>

        {/* RIGHT: MENU */}
        <div className="flex gap-6 text-sm font-medium text-gray-700">
          <Link href="/" className="hover:text-primary transition">Home</Link>
          <Link href="/technology" className="hover:text-primary transition">Technology</Link>
          <Link href="/poss" className="hover:text-primary transition">PoSS</Link>
          <Link href="/curators" className="hover:text-primary transition">Curators</Link>
          <Link href="/genesis" className="hover:text-primary transition">Genesis</Link>
          <Link href="/roadmap" className="hover:text-primary transition">Roadmap</Link>
          <Link href="/docs" className="hover:text-primary transition">Docs</Link>
          <Link href="/legal" className="hover:text-primary transition">Legal</Link>
        </div>

      </div>
    </nav>
  );
}

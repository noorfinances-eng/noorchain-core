"use client";

import Link from "next/link";

export default function NavBar() {
  return (
    <nav className="w-full border-b bg-white">
      <div className="container flex items-center justify-between py-4">
        <div className="text-xl font-bold tracking-tight">
          NOORCHAIN
        </div>

        <div className="flex gap-5 text-sm text-gray-700">
          <Link href="/" className="hover:text-black">
            Home
          </Link>
          <Link href="/technology" className="hover:text-black">
            Technology
          </Link>
          <Link href="/poss" className="hover:text-black">
            PoSS
          </Link>
          <Link href="/curators" className="hover:text-black">
            Curators
          </Link>
          <Link href="/genesis" className="hover:text-black">
            Genesis
          </Link>
          <Link href="/roadmap" className="hover:text-black">
            Roadmap
          </Link>
          <Link href="/docs" className="hover:text-black">
            Docs
          </Link>
          <Link href="/legal" className="hover:text-black">
            Legal
          </Link>
        </div>
      </div>
    </nav>
  );
}

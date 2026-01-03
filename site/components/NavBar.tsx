"use client";

import Link from "next/link";
import Image from "next/image";
import { useState } from "react";

export default function NavBar() {
  const [open, setOpen] = useState(false);

    const navItems = [
    { href: "/#top", label: "Home" },
    { href: "/#technology", label: "Technology" },
    { href: "/#poss", label: "PoSS" },
    { href: "/#curators", label: "Curators" },
    { href: "/#genesis", label: "Genesis" },
    { href: "/#roadmap", label: "Roadmap" },
    { href: "/#docs", label: "Docs" },
    { href: "/#legal", label: "Legal" },
  ];

  return (
    <header className="sticky top-0 z-50 w-full bg-white border-b border-gray-soft shadow-sm">
      <div className="container flex items-center justify-between py-4">
        
        {/* LEFT: LOGO + NAME */}
        <Link href="/#top" className="flex items-center gap-3">
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

        {/* DESKTOP MENU */}
        <nav className="hidden md:flex gap-6 text-sm font-medium text-gray-700">
          {navItems.map((item) => (
            <Link 
              key={item.href} 
              href={item.href} 
              className="hover:text-primary transition"
            >
              {item.label}
            </Link>
          ))}
        </nav>

        {/* MOBILE HAMBURGER BUTTON */}
        <button
          onClick={() => setOpen(!open)}
          className="md:hidden p-2 rounded-md text-gray-700 hover:bg-gray-100 focus:outline-none"
          aria-label="Open menu"
        >
          <div className="space-y-1">
            <span
              className={`block h-0.5 w-6 bg-current transition ${
                open ? "rotate-45 translate-y-1" : ""
              }`}
            />
            <span
              className={`block h-0.5 w-6 bg-current transition ${
                open ? "opacity-0" : "opacity-100"
              }`}
            />
            <span
              className={`block h-0.5 w-6 bg-current transition ${
                open ? "-rotate-45 -translate-y-1" : ""
              }`}
            />
          </div>
        </button>
      </div>

      {/* MOBILE MENU DROPDOWN */}
      {open && (
        <nav className="md:hidden bg-white border-t border-gray-soft shadow-sm">
          <div className="container flex flex-col py-3">
            {navItems.map((item) => (
              <Link
                key={item.href}
                href={item.href}
                className="py-2 text-sm font-medium text-gray-800 hover:text-primary transition"
                onClick={() => setOpen(false)}
              >
                {item.label}
              </Link>
            ))}
          </div>
        </nav>
      )}
    </header>
  );
}

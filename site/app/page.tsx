"use client";

import Image from "next/image";
import { useEffect, useRef } from "react";
import Reveal from "../components/ui/Reveal";

/**
 * Home — Refonte A+ (Hero + Signal Board)
 * - No external deps
 * - No randomness (no hydration mismatch)
 * - Subtle parallax via CSS vars (mouse), disabled for prefers-reduced-motion
 */

const SIGNAL_PARTICLES = [
  { left: "10%", top: "22%", s: 10, o: 0.18, d: "18s", delay: "-3s" },
  { left: "14%", top: "66%", s: 14, o: 0.12, d: "22s", delay: "-9s" },
  { left: "22%", top: "38%", s: 8, o: 0.16, d: "16s", delay: "-6s" },
  { left: "32%", top: "18%", s: 9, o: 0.14, d: "20s", delay: "-11s" },
  { left: "36%", top: "72%", s: 12, o: 0.11, d: "24s", delay: "-7s" },
  { left: "48%", top: "28%", s: 13, o: 0.10, d: "26s", delay: "-12s" },
  { left: "56%", top: "62%", s: 9, o: 0.12, d: "19s", delay: "-5s" },
  { left: "64%", top: "20%", s: 8, o: 0.13, d: "21s", delay: "-10s" },
  { left: "72%", top: "46%", s: 11, o: 0.10, d: "27s", delay: "-14s" },
  { left: "82%", top: "30%", s: 12, o: 0.09, d: "25s", delay: "-8s" },
  { left: "86%", top: "70%", s: 14, o: 0.08, d: "30s", delay: "-15s" },
];

function PoSSSignalBoard() {
  return (
    <div className="relative overflow-hidden rounded-2xl border border-white/15 bg-white/[0.06] shadow-[0_26px_80px_rgba(0,0,0,0.28)] backdrop-blur-md">
      {/* top sheen */}
      <div className="pointer-events-none absolute inset-0 bg-[radial-gradient(800px_circle_at_20%_10%,rgba(255,255,255,0.18),transparent_55%),radial-gradient(700px_circle_at_85%_30%,rgba(255,255,255,0.10),transparent_60%)]" />

      {/* animated grid stream */}
      <div className="pointer-events-none absolute inset-0 opacity-[0.28] bg-[linear-gradient(to_right,rgba(255,255,255,0.22)_1px,transparent_1px),linear-gradient(to_bottom,rgba(255,255,255,0.22)_1px,transparent_1px)] bg-[size:42px_42px] noor-grid-stream" />

      {/* content */}
      <div className="relative p-6">
        <div className="flex items-start justify-between gap-4 mb-5">

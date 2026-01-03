"use client";

import React from "react";

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
] as const;

export default function HeroGlow() {
  return (
    <div className="pointer-events-none absolute inset-0 overflow-hidden -z-10">
      {/* Base gradient (match HERO) */}
      <div className="absolute inset-0 bg-gradient-to-br from-[#1A6AFF] to-[#00D1B2]" />

      {/* Spotlight + depth (reuse hero CSS class) */}
      <div className="absolute inset-0 noor-spotlight" />

      {/* Subtle grid (match HERO) */}
      <div className="absolute inset-0 opacity-[0.12] bg-[linear-gradient(to_right,rgba(255,255,255,0.35)_1px,transparent_1px),linear-gradient(to_bottom,rgba(255,255,255,0.35)_1px,transparent_1px)] bg-[size:56px_56px]" />

      {/* Particles (reuse hero particle class/vars) */}
      <div className="absolute inset-0 noor-parallax-1">
        {SIGNAL_PARTICLES.map((p, i) => (
          <span
            key={i}
            className="noor-particle"
            style={
              {
                left: p.left,
                top: p.top,
                width: `${p.s}px`,
                height: `${p.s}px`,
                opacity: p.o,
                ["--dur" as any]: p.d,
                ["--delay" as any]: p.delay,
              } as React.CSSProperties
            }
          />
        ))}
      </div>

      {/* Scan line (reuse hero scan class) */}
          </div>
  );
}

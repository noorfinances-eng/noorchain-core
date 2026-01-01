import React from "react";

export function PageHeader({
  label,
  title,
  intro,
}: {
  label: string;
  title: string;
  intro: React.ReactNode;
}) {
  return (
    <>
      {/* LABEL */}
      <div className="inline-flex items-center gap-2 rounded-full bg-white border border-gray-soft px-3 py-1 mb-6">
        <span className="h-2 w-2 rounded-full bg-primary" />
        <span className="text-xs font-medium uppercase tracking-wide text-gray-700">
          {label}
        </span>
      </div>

      {/* TITLE */}
      <h1 className="text-3xl md:text-4xl font-extrabold tracking-tight text-navy mb-4">
        {title}
      </h1>

      {/* INTRO */}
      <p className="text-lg text-gray-700 leading-relaxed mb-10 border-l-4 border-primary pl-4 bg-white/60 py-3 rounded-r-lg">
        {intro}
      </p>
    </>
  );
}

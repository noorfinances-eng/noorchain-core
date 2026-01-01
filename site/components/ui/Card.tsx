import React from "react";

export function Card({
  children,
  className = "",
}: {
  children: React.ReactNode;
  className?: string;
}) {
  return (
    <section
      className={[
        "p-6 border border-gray-soft rounded-xl bg-white shadow-sm",
        "transition-all duration-200 hover:shadow-md hover:-translate-y-0.5",
        className,
      ].join(" ")}
    >
      {children}
    </section>
  );
}

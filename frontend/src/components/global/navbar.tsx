'use client';
import { useState } from 'react';
import Link from 'next/link';
import Image from 'next/image';

type NavbarProps = {
  role: 'kandidat' | 'perekrut';
};

const Navbar = ({ role }: NavbarProps) => {
  const [isOpen, setIsOpen] = useState(false);

  const navItems = role === 'kandidat'
    ? [
        { label: 'loker', href: '/kandidat' },
        { label: 'status aplikasi', href: '/kandidat/status-aplikasi' },
        { label: 'profil', href: '/kandidat/profile' },
        { label: 'logout', href: '/' },
      ]
    : [
        { label: 'loker', href: '/perekrut' },
        { label: 'profil', href: '/perekrut/profile' },
        { label: 'logout', href: '/' },
      ];

  return (
    <header className="border-b">
      <div className="container mx-auto py-6 flex items-center justify-between">
        {/* Logo */}
        <div className="flex items-center">
          <Link href="/">
            <Image
              src="/LokerHub_Logo.svg"
              alt="logo"
              width={222}
              height={42}
              className="w-full h-auto"
              priority
            />
          </Link>
        </div>

        {/* Desktop Menu */}
        <div className="hidden md:flex items-center gap-8">
          {navItems.map((item, index) => (
            <Link
              key={index}
              href={item.href}
              className="text-lg text-[#1e3a5f] hover:bold font-medium"
            >
              {item.label}
            </Link>
          ))}
        </div>

        {/* Hamburger Icon */}
        <div className="md:hidden">
          <button
            onClick={() => setIsOpen(!isOpen)}
            className="text-[#1e3a5f] focus:outline-none cursor-pointer"
          >
            <svg
              className="h-6 w-6"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              {isOpen ? (
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M6 18L18 6M6 6l12 12"
                />
              ) : (
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M4 8h16M4 16h16"
                />
              )}
            </svg>
          </button>
        </div>
      </div>

      {/* Dropdown Menu */}
      {isOpen && (
        <div className="md:hidden px-6 pb-4">
          <div className="flex flex-col gap-2">
            {navItems.map((item, index) => (
              <Link
                key={index}
                href={item.href}
                className="text-sm text-[#1e3a5f] py-1 hover:underline font-medium"
              >
                {item.label}
              </Link>
            ))}
          </div>
        </div>
      )}
    </header>
  );
};

export default Navbar;

"use client"

import type React from "react"
import Link from "next/link"
import Image from 'next/image';
import { Button } from '@/components/ui/button';
import { GoX } from "react-icons/go";
import { useRouter } from "next/navigation"

export default function Login() {
  const router = useRouter()
    
  return(
    <div className="min-h-screen flex items-center justify-center bg-[#f9f5f0] p-2">
      <div className="w-full max-w-5xl grid grid-cols-1 md:grid-cols-2 gap-8 items-center">
        {/* Logo Section */}
        <div className="flex justify-center items-center w-full mx-auto">
          <div className="w-96 h-32 relative items-center justify-center">
            <Image
              src="/LokerHub_Logo.svg"
              alt="logo"
              fill
              priority
              className="object-contain"
              sizes="(max-width: 768px) 10vw, 50vw"/>
          </div>
        </div>
                
        {/* Login Form Section */}
          <div className="relative bg-white rounded-lg p-10 max-w-md mx-auto shadow-lg w-full">
          {/* Close Button */}
            <Button 
              variant="ghost" 
              className="absolute right-6 top-6 hover:text-gray-700 cursor-pointer" 
              onClick={() => router.push("/")}>
              <GoX className="h-6 w-6" />
            </Button>
                    
            <h2 className="text-2xl font-semibold text-center mb-8 text-[#1e3a5f]">Masuk Ke Akunmu</h2>
                    
            <form className="space-y-5">
              <div>
                <input
                  type="email"
                  placeholder="Email"
                  required
                  className="w-full px-3 py-2 text-lg border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                />
              </div>
                        
              <div>
                <input
                  type="password"
                  placeholder="Password"
                  required
                  className="w-full px-3 py-2 text-lg border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                />
              </div>
                        
              <div className="text-center text-base flex justify-center gap-2">
                <p>Belum Punya Akun?</p>
                <Link href="/auth/register" className="text-[#1e3a5f] font-medium hover:underline">
                  Buat Akun
                </Link>
              </div>          
                        
              <Button 
                type="submit"
                className="w-full py-2 text-lg bg-[#a8c5e2] hover:bg-[#97b8d9] text-[#1e3a5f] font-medium rounded-md transition-colors cursor-pointer">
                Masuk
              </Button>
            </form>
        </div>
      </div>
    </div>
  )
}
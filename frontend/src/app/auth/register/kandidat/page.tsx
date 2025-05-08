"use client"

import type React from "react"
import Link from "next/link"
import Image from 'next/image';
import { Button } from '@/components/ui/button';
import { GoX } from "react-icons/go";
import { useRouter } from "next/navigation"
import { FiArrowLeft } from "react-icons/fi";


export default function RegisterKandidat() {
    const router = useRouter()

    return(
        <div className="min-h-screen flex items-center justify-center bg-[#f9f5f0] p-2">
            <div className="w-full max-w-5xl grid grid-cols-1 md:grid-cols-2 gap-8 items-center">
                {/* Logo Section */}
                <div className="flex justify-center items-center w-full mx-auto">
                    <div className="w-96 h-32 relative items-center justofy-center">
                        <Image
                            src="/LokerHub_Logo.svg"
                            alt="logo"
                            fill
                            priority
                            className="object-contain"
                            sizes="(max-width: 768px) 10vw, 50vw"
                        />
                    </div>
                    
                </div>
                
                {/* Login Form Section */}
                <div className="relative bg-white rounded-lg p-10 max-w-md mx-auto shadow-lg w-full">
                    {/* Close Button */}
                    <Button variant="ghost" className="absolute right-6 top-6 hover:text-gray-700 cursor-pointer" onClick={() => router.push("/")}>
                        <GoX className="h-6 w-6" />
                    </Button>
                    
                    <div className="flex items-center gap-3 mb-6">
                        <button onClick={() => router.push("/auth/register")} className="text-gray-500 hover:text-gray-700">
                            <FiArrowLeft className="h-5 w-5 cursor-pointer" />
                        </button>
                        <h2 className="text-xl font-semibold text-[#1e3a5f]">Registrasi Kandidat</h2>
                    </div>
                    
                    <form className="space-y-5">
                        <div>
                            <input
                            type="email"
                            name="email"
                            placeholder="Email"
                            required
                            className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                            />
                        </div>

                        <div>
                            <input
                            type="password"
                            name="password"
                            placeholder="Password"
                            required
                            className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                            />
                        </div>

                        <div>
                            <input
                            type="password"
                            name="confirmPassword"
                            placeholder="Re-Enter Password"
                            required
                            className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                            />
                        </div>

                        <button
                            type="submit"
                            className="w-full py-2 bg-[#a8c5e2] hover:bg-[#97b8d9] text-[#1e3a5f] font-medium rounded-md transition-colors cursor-pointer"
                        >
                            Buat Akun
                        </button>
                        </form>
                </div>
            </div>
        </div>
    )
}
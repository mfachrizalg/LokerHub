"use client"

import Image from "next/image"
import Link from "next/link"
import { useRouter } from "next/navigation"
import React from "react"
import Navbar from "@/components/global/navbar"
import Footer from "@/components/global/footer"

const profile = {
  fullName: "Aditya Pratama",
  title: "Senior HR in Shopee Indonesia",
  photoUrl: "/placeholder.svg?height=120&width=120",
  phoneNumber: "081548797755",
  company: {
    name: "Shopee",
    industry: "E-Commerce",
    location: "Jakarta Pusat",
    logoUrl: "/placeholder.svg?height=80&width=80",
  },
}

export default function ProfileRekruter() {
  const router = useRouter()

  return (
    <div className="min-h-screen flex flex-col bg-white">
      {/* Header */}
      <Navbar role="perekrut"/>

      {/* Profile section */}
      <div className="container mx-auto px-4 py-4 mb-10">
        <h1 className="text-2xl font-bold text-center text-[#1e3a5f] mb-4">Profil</h1>

        <div className="max-w-2xl mx-auto space-y-2">
          {/* Recruiter Profile Section */}
          <div className="container mx-auto px-4 py-4">
            <div className="max-w-2xl mx-auto">
              <div className="bg-[#f8f0e5] rounded-lg border border-[#e8d5c4] py-3">
               <h3 className="text-lg font-bold text-center text-[#1e3a5f] mb-2">Profil Rekruter</h3>
                <div className="flex flex-col items-center mb-4">
                  <div className="relative h-32 w-32 rounded-full border border-black overflow-hidden mb-4">
                    <Image 
                      src={profile.photoUrl || "/placeholder.svg"} 
                      alt={profile.fullName}
                      fill
                      className="object-cover"
                    />
                  </div>
                  <h2 className="text-xl font-bold text-[#1e3a5f]">{profile.fullName}</h2>
                  <p className="text-lg text-center text-gray-600">{profile.title}</p>
                </div>

                <div className="px-6 text-center">
                  <div>
                    <p className="text-gray-500">Nomor Telepon</p>
                    <p className="font-medium">{profile.phoneNumber || "-"}</p>
                  </div>
                  
                </div>  
              </div>
            </div>
          </div>

          {/* Company Profile Section */}
          <div className="container mx-auto px-4 py-2">
            <div className="max-w-2xl mx-auto">
              <div className="bg-[#f8f0e5] rounded-lg border border-[#e8d5c4] py-3">
                <h3 className="text-lg font-bold text-center text-[#1e3a5f] mb-2">Profil Perusahaan</h3>

                <div className="flex flex-col items-center mb-4">
                  <div className="relative h-24 w-24 bg-red-100 border border-black rounded-md overflow-hidden mb-4 flex items-center justify-center">
                    <Image 
                      src={profile.company.logoUrl || "/placeholder.svg"} 
                      alt={profile.company.name}
                      fill
                      className="object-contain p-2"
                    />
                  </div>
                  <h2 className="text-xl font-bold text-[#1e3a5f]">{profile.company.name}</h2>
                  <p className="text-lg text-center text-gray-600">{profile.company.industry}</p>
                </div>

                <div className="px-6 text-center">
                  <div>
                    <p className="text-gray-500">Lokasi</p>
                    <p className="font-medium">{profile.company.location || "-"}</p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          {/* Edit Profile Button */}
          <div className="flex justify-center mt-6">
            <Link
              href="/perekrut/profile/edit"
              className="bg-[#1e3a5f] text-white px-6 py-2 rounded-md hover:bg-[#1e3a5f]/90 transition-colors disabled:opacity-50 cursor-pointer"
            >
              ubah profil
            </Link>
          </div>
        </div>
      </div>

      {/* Footer */}
      <Footer />
    </div>
  )
}
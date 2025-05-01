'use client'

import Image from "next/image"
import { useRouter } from "next/navigation"
import { Button } from "@/components/ui/button"
import Navbar from "@/components/global/navbar"
import Footer from "@/components/global/footer"

const profile = {
  fullName: "Aditya Pratama",
  title: "Fresh Graduate Teknologi Informasi Universitas Gadjah Mada",
  photoUrl: "/placeholder.svg?height=120&width=120",
  phoneNumber: "081548797755",
  education: "S1",
  location: "Yogyakarta",
  cvUrl: "cv.pdf",
  fieldOfInterest: "IT"
}

export default function ProfileKandidat() {
  const router = useRouter()

  const handleEditProfile = () => {
    router.push('profile/edit')
  }

  return (
    <div className="min-h-screen flex flex-col bg-white">
      {/* Header */}
      <Navbar role="kandidat" />

      {/* Profile Section */}
      <div className="container mx-auto px-4 py-8 mb-10">
        <h1 className="text-2xl font-bold text-center text-[#1e3a5f] mb-8">Profil</h1>
        <div className="max-w-2xl mx-auto">
          <div className="bg-[#f8f0e5] rounded-lg border border-[#e8d5c4] py-6 px-4">
            <div className="flex flex-col items-center mb-6">
              <div className="relative h-32 w-32 rounded-full border border-gray-200 overflow-hidden mb-4">
                <Image 
                  src={profile.photoUrl} 
                  alt={profile.fullName}
                  fill
                  className="object-cover"
                />
              </div>
              <h2 className="text-xl font-bold text-[#1e3a5f]">{profile.fullName}</h2>
              <p className="text-lg text-center text-gray-600">{profile.title}</p>
            </div>
            
            <div className="grid grid-cols-2 gap-4 text-center text-sm">
              <div>
                <p className="text-gray-500">Nomor Telepon</p>
                <p className="font-medium">{profile.phoneNumber}</p>
              </div>
              <div>
                <p className="text-gray-500">Pendidikan Terakhir</p>
                <p className="font-medium">{profile.education}</p>
              </div>
              <div>
                <p className="text-gray-500">Lokasi</p>
                <p className="font-medium">{profile.location}</p>
              </div>
              <div>
                <p className="text-gray-500">CV</p>
                <a 
                  href={profile.cvUrl} 
                  target="_blank" 
                  rel="noopener noreferrer"
                  className="inline-block border border-[#1e3a5f] text-[#1e3a5f] px-3 py-1 rounded-md hover:bg-[#1e3a5f]/10"
                >
                  {profile.cvUrl}
                </a>
              </div>
            </div>  
          </div>

          <div className="flex justify-center pt-10">
            <Button
              onClick={handleEditProfile}
              className="bg-[#1e3a5f] text-white px-6 py-2 rounded-md hover:bg-[#1e3a5f]/90 transition-colors disabled:opacity-50 cursor-pointer"
            >
              Ubah profil
            </Button>
          </div>
        </div>
      </div>

      {/* Footer */}
      <Footer />
    </div>
  )
}

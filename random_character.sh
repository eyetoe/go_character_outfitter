#!/bin/bash

################################################################################
CLOSE_RANGE=(
             Knife
             Sledgehammer
             Chainsaw
             Welding_Torch
             Ice_Pick
             Crowbar
             Wrench
             Broom
             Stun_Stick
             Fart
             Pokey_Stick
            ) 
CLOSE_CHOICE=$(shuf --input-range=0-$(( ${#CLOSE_RANGE[*]} - 1 )) | head -n 1)
THIS_CLOSE_CHOICE="${CLOSE_RANGE[$CLOSE_CHOICE]}"

################################################################################
MEDIUM_RANGE=(
              Flame_Thrower
              Shotgun
              Double_Barrel_Shotgun
              Gernade_Launcher
              Throwing_Knife
              Stun_Gun
              Short_Bow
              Crossbow_Pistol
             )
MEDIUM_CHOICE=$(shuf --input-range=0-$(( ${#MEDIUM_RANGE[*]} - 1 )) | head -n 1)
THIS_MEDIUM_CHOICE="${MEDIUM_RANGE[$MEDIUM_CHOICE]}"

################################################################################
GERNADE=(
              Explosive
              Flash_Bang
              Tear_Gas
              Bouncy
              Sticky
              Incindiary
              Cryo
              Tesla
              Molotov
              Cluster
              Shrapnel
             )
GERNADE_CHOICE=$(shuf --input-range=0-$(( ${#GERNADE[*]} - 1 )) | head -n 1)
THIS_GERNADE_CHOICE="${GERNADE[$GERNADE_CHOICE]}"

################################################################################
LONG_RANGE=(
              Pistol
              SMG
              Assault_Rifle
              Hunting_Rifle
              Sniper_Rifle
              Rocker_Launcher
              Crossbow
              Long_Bow
             )
LONG_CHOICE=$(shuf --input-range=0-$(( ${#LONG_RANGE[*]} - 1 )) | head -n 1)
THIS_LONG_CHOICE="${LONG_RANGE[$LONG_CHOICE]}"

################################################################################
FIRST_NAME=(
  Wap
  Dink
  Woosh
  Bark
  Slim
  Walnut
  Dingo
  Poodle
  Sam
  Flip
  Sherbert
  Spatula
  Fecalpap
  Smorgasbord
  Silly
  Joe
  Ted
  Burn
  Tingle
  Splash
  Johnny
  Tuna
)
FIRST_CHOICE=$(shuf --input-range=0-$(( ${#FIRST_NAME[*]} - 1 )) | head -n 1)
THIS_FIRST_CHOICE="${FIRST_NAME[$FIRST_CHOICE]}"

LAST_NAME=(
  Waperson
  Dinkski
  Wooshblatz
  Barker
  Slime
  Cashew
  Stink
  Smith
  Curry
  Palendrome
  Fooperson
  Happyslap
  Koala-nose
  Slothface
  Wilikins
  Wilson
  Jammypants
  Googlethat
  Nevermind
  Tuner
  Fishface
)
LAST_CHOICE=$(shuf --input-range=0-$(( ${#LAST_NAME[*]} - 1 )) | head -n 1)
THIS_LAST_CHOICE="${LAST_NAME[$LAST_CHOICE]}"

echo -e "\n!! YOUR CHARACTER !!\n"

echo -e "\nNAME: $THIS_FIRST_CHOICE $THIS_LAST_CHOICE\n"

echo -e "\nWEAPONS:\n"
echo -e "CLOSE: $THIS_CLOSE_CHOICE"
echo -e "MEDIUM: $THIS_MEDIUM_CHOICE"
echo -e "LONG: $THIS_LONG_CHOICE"
echo -e "GERNADE: $THIS_GERNADE_CHOICE"
echo -e "\n\n\n\n"

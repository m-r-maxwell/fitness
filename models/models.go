package models

type Exercise struct {
	Name          string
	MuscleGroup   string
	Reps          int
	WorkingTime   int
	Sets          int
	IsAdvaned     bool
	Modifications []string
	Description   string
}

var PushUps = Exercise{
	Name:          "Push Ups",
	MuscleGroup:   "Chest",
	Reps:          10,
	WorkingTime:   0,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Knee Push Ups", "Incline Push Ups"},
	Description:   "Push ups are a great exercise for your chest, triceps, and shoulders. Keep your core tight and your body in a straight line from head to heels. Lower your body until your chest nearly touches the floor, then push yourself back up to the starting position.",
}

var DeclinePushUps = Exercise{
	Name:          "Decline Push Ups",
	MuscleGroup:   "Chest",
	Reps:          6,
	WorkingTime:   0,
	Sets:          3,
	IsAdvaned:     true,
	Modifications: []string{},
	Description:   "Decline push ups are a more advanced variation of the standard push up. Elevating your feet increases the difficulty of the exercise and places more emphasis on the upper chest and shoulders. Keep your core tight and your body in a straight line from head to heels. Lower your body until your chest nearly touches the floor, then push yourself back up to the starting position.",
}

var Squats = Exercise{
	Name:          "Squats",
	MuscleGroup:   "Legs",
	Reps:          15,
	WorkingTime:   0,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Chair Squats", "Pulse Squats", "Sumo Squats", "Seal Jacks"},
	Description:   "Squats are a great exercise for your quads, hamstrings, and glutes. Stand with your feet shoulder-width apart and your toes pointed slightly outward. Lower your body by pushing your hips back and bending your knees. Keep your chest up and your back straight. Lower yourself until your thighs are parallel to the ground, then push yourself back up to the starting position.",
}

var Lunges = Exercise{
	Name:          "Lunges",
	MuscleGroup:   "Legs",
	Reps:          10,
	WorkingTime:   0,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Reverse Lunges", "Side Lunges", "Curtsy Lunges"},
	Description:   "Lunges are a great exercise for your quads, hamstrings, and glutes. Stand with your feet hip-width apart and your hands on your hips. Take a big step forward with your right foot and lower your body until your right thigh is parallel to the ground. Push yourself back up to the starting position and repeat on the other side.",
}

var Plank = Exercise{
	Name:          "Plank",
	MuscleGroup:   "Core",
	Reps:          0,
	WorkingTime:   30,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Forearm Plank", "Side Plank", "Plank with Leg Lift"},
	Description:   "Planks are a great exercise for your core, shoulders, and back. Start in a push up position with your hands directly under your shoulders. Lower yourself onto your forearms and hold your body in a straight line from head to heels. Engage your core and hold the position for the prescribed time.",
}

var Crunches = Exercise{
	Name:          "Crunches",
	MuscleGroup:   "Core",
	Reps:          15,
	WorkingTime:   0,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Bicycle Crunches", "Russian Twists", "Leg Raises"},
	Description:   "Crunches are a great exercise for your abs. Lie on your back with your knees bent and your feet flat on the floor. Place your hands behind your head and lift your shoulders off the ground. Engage your core and crunch your upper body towards your knees. Lower yourself back down and repeat for the prescribed number of reps.",
}

var JumpingJacks = Exercise{
	Name:          "Jumping Jacks",
	MuscleGroup:   "Full Body",
	Reps:          0,
	WorkingTime:   30,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Seal Jacks", "Star Jumps", "Squat Jacks", "Arm Raises and Leg Behind Other Leg"},
	Description:   "Jumping jacks are a great exercise for your heart and lungs. Start with your feet together and your arms at your sides. Jump up and spread your legs while raising your arms above your head. Jump back to the starting position and repeat for the prescribed time.",
}

var MountainClimbers = Exercise{
	Name:          "Mountain Climbers",
	MuscleGroup:   "Full Body",
	Reps:          0,
	WorkingTime:   30,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Cross Body Mountain Climbers", "Plank Jacks", "Plank to Push Up"},
	Description:   "Mountain climbers are a great exercise for your core, shoulders, and legs. Start in a push up position with your hands directly under your shoulders. Drive your right knee towards your chest, then quickly switch legs. Keep your core tight and your body in a straight line from head to heels. Repeat for the prescribed time.",
}

var DownwardDog = Exercise{
	Name:          "Downward Dog",
	MuscleGroup:   "Full Body",
	Reps:          0,
	WorkingTime:   30,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Puppy Pose", "Dolphin Pose", "Ragdoll Pose"},
	Description:   "Downward dog is a great yoga pose that stretches your hamstrings, calves, and shoulders. Start on your hands and knees with your wrists under your shoulders and your knees under your hips. Lift your hips up and back, straightening your arms and legs. Press your heels towards the ground and relax your head between your arms.",
}

var ChildsPose = Exercise{
	Name:          "Child's Pose",
	MuscleGroup:   "Full Body",
	Reps:          0,
	WorkingTime:   30,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Extended Child's Pose", "Thread the Needle", "Puppy Pose"},
	Description:   "Child's pose is a great yoga pose that stretches your back, hips, and thighs. Start on your hands and knees with your wrists under your shoulders and your knees under your hips. Sit back on your heels and lower your chest towards the ground. Extend your arms out in front of you and relax your forehead on the mat.",
}

var CactusPose = Exercise{
	Name:          "Cactus Pose",
	MuscleGroup:   "Full Body",
	Reps:          0,
	WorkingTime:   30,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Reverse Prayer Pose", "Cow Face Pose", "Eagle Arms"},
	Description:   "Cactus pose is a great yoga pose that stretches your chest, shoulders, and arms. Stand with your feet hip-width apart and your arms out to the sides at shoulder height. Bend your elbows and bring your forearms together in front of you. Press your palms together and lift your elbows up to shoulder height.",
}

var BoundAnglePose = Exercise{
	Name:          "Bound Angle Pose",
	MuscleGroup:   "Full Body",
	Reps:          0,
	WorkingTime:   30,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Reclining Bound Angle Pose", "Butterfly Pose", "Wide-Legged Forward Fold"},
	Description:   "Bound angle pose is a great yoga pose that stretches your inner thighs, groins, and knees. Sit on the floor with your legs extended in front of you. Bend your knees and bring the soles of your feet together. Let your knees fall out to the sides and bring your heels as close to your pelvis as you can.",
}

var GluteBridge = Exercise{
	Name:          "Glute Bridge",
	MuscleGroup:   "Glutes",
	Reps:          15,
	WorkingTime:   0,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Single Leg Glute Bridge", "Marching Glute Bridge", "Hip Thrusts"},
	Description:   "Glute bridge is a great exercise for your glutes and hamstrings. Lie on your back with your knees bent and your feet flat on the floor. Press through your heels and lift your hips up towards the ceiling. Squeeze your glutes at the top of the movement, then lower yourself back down and repeat for the prescribed number of reps.",
}

var Superman = Exercise{
	Name:          "Superman",
	MuscleGroup:   "Back",
	Reps:          15,
	WorkingTime:   0,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Bird Dog", "Swimmers", "Reverse Snow Angels"},
	Description:   "Superman is a great exercise for your lower back and glutes. Lie on your stomach with your arms extended in front of you and your legs straight. Lift your arms, chest, and legs off the ground at the same time. Hold the position for a second, then lower yourself back down and repeat for the prescribed number of reps.",
}

var CatCow = Exercise{
	Name:          "Cat Cow",
	MuscleGroup:   "Back",
	Reps:          0,
	WorkingTime:   30,
	Sets:          2,
	IsAdvaned:     false,
	Modifications: []string{"Thread the Needle", "Child's Pose", "Downward Dog"},
	Description:   "Cat cow is a great yoga pose that stretches your back and shoulders. Start on your hands and knees with your wrists under your shoulders and your knees under your hips. Inhale and arch your back, lifting your head and tailbone towards the ceiling. Exhale and round your back, tucking your chin to your chest and pressing your hands into the ground.",
}

var BirdDog = Exercise{
	Name:          "Bird Dog",
	MuscleGroup:   "Back",
	Reps:          10,
	WorkingTime:   0,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Bird Dog with Leg Lift", "Bird Dog with Arm Extension", "Bird Dog with Knee to Elbow"},
	Description:   "Bird dog is a great exercise for your core and lower back. Start on your hands and knees with your wrists under your shoulders and your knees under your hips. Extend your right arm and left leg out in front of you, keeping your back flat and your core engaged. Hold the position for a second, then lower yourself back down and repeat on the other side.",
}

var BackExtension = Exercise{
	Name:          "Back Extension",
	MuscleGroup:   "Back",
	Reps:          15,
	WorkingTime:   0,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Superman", "Bird Dog", "Swimmers"},
	Description:   "Back extension is a great exercise for your lower back and glutes. Lie face down on the ground with your arms extended in front of you. Lift your chest and legs off the ground at the same time, squeezing your glutes at the top of the movement. Hold the position for a second, then lower yourself back down and repeat for the prescribed number of reps.",
}

var HighKnees = Exercise{
	Name:          "High Knees",
	MuscleGroup:   "Full Body",
	Reps:          0,
	WorkingTime:   30,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Butt Kicks", "Running in Place", "Jump Rope"},
	Description:   "High knees are a great exercise for your heart and lungs. Start with your feet hip-width apart and your arms at your sides. Run in place, bringing your knees up towards your chest as high as you can. Pump your arms as if you were running and keep up the pace for the prescribed time.",
}

var StandingSideCrunch = Exercise{
	Name:          "Standing Side Crunch",
	MuscleGroup:   "Core",
	Reps:          20,
	WorkingTime:   0,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Standing Oblique Crunch", "Standing Knee to Elbow", "Standing Side Bend"},
	Description:   "Standing side crunch is a great exercise for your obliques. Stand with your feet hip-width apart and your hands behind your head. Lift your right knee towards your right elbow, crunching your obliques. Lower your leg back down and repeat on the other side.",
}

var StandingAlternatingCrunch = Exercise{
	Name:          "Standing Alternating Crunch",
	MuscleGroup:   "Core",
	Reps:          20,
	WorkingTime:   0,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Standing Knee to Elbow", "Standing Side Crunch", "Standing Oblique Crunch"},
	Description:   "Standing alternating crunch is a great exercise for your core. Stand with your feet hip-width apart and your hands behind your head. Lift your right knee towards your left elbow, crunching your abs. Lower your leg back down and repeat on the other side.",
}

var StandingToeTouches = Exercise{
	Name:          "Standing Toe Touches",
	MuscleGroup:   "Core",
	Reps:          20,
	WorkingTime:   0,
	Sets:          3,
	IsAdvaned:     false,
	Modifications: []string{"Standing Knee to Elbow", "Standing Side Crunch", "Standing Alternating Crunch"},
	Description:   "Standing toe touches are a great exercise for your core and hamstrings. Stand with your feet hip-width apart and your arms extended overhead. Bend at the waist and reach your right hand towards your left foot. Return to the starting position and repeat on the other side.",
}

var List = []Exercise{
	PushUps,
	DeclinePushUps,
	Squats,
	Lunges,
	Plank,
	Crunches,
	JumpingJacks,
	MountainClimbers,
	DownwardDog,
	ChildsPose,
	CactusPose,
	BoundAnglePose,
	GluteBridge,
	Superman,
	CatCow,
	BirdDog,
	BackExtension,
	HighKnees,
	StandingSideCrunch,
	StandingAlternatingCrunch,
	StandingToeTouches,
}
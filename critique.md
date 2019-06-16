# Critique
So, I'm following this guy's tutorial because I want to see how he does things, and get some exposure to golang through a domain that I know reasonably well.  The video series began episode 1 with some motivation for why the author thinks go is suitable to this kind of thing, but as I work through the tutorial I'm developing the opinion that it really isn't.  So this markdown page is my way of thinking through why that is exactly, in the hopes of learning something. 

## Why golang?
The tutorial author puts forth that go is a nicer language to work in than some of the alternatives.  For example they point out that C++ is a very popular language for games development, and contrast the ease of doing things in Go versus that.  To me, this doesn't consider the better alternatives.  Something like Unity lets you work in C#, which to me is vastly easier to work with.  In golang, I see that you're still having to pay attention to the differences between pointers and not, and that just seems too low level to me.  Additionally, it might just be the tutorial author's style but this all seems very verbose.  Examples of verbosity: 
* Having to check error after every function call.  While this enforces correctness, exceptions work fine. 
* The interface defined for components can't implement default behavior, so every component has to implement every possible behaviour.  This gets very repetitve. 

## Why do a custom engine at all? 
Unity has a huge user base and support all over the internet, is free until you start selling your game (basically), and has the asset store to find things.  If you go from prototyping a game to actually making it, you're going to find way more people who can contribute to your game with the Unity engine.  For me in particular, if I were to actually make something I would want to have a technical artist type of person to assist and they would most likely be more productive with Unity than with some custom engine. 

## Specific programming criticisms 

### Collision is all screwy
I think adding collision as a base attribute of all entities in the ECS was a bad call.  All entities have a circle collider, so that's less flexible than having different components that allow e.g. BoxCollider, CircleCollider, etc.  Also, the colliders have their own position which always tracks the position of the entity, so why is it a separate value that has to be updated each frame instead of just referencing that position?  Another thing, the initial implementation in the video screwed up not colliding with unrelated entities; e.g. the player. 

### Update time step
This game engine does the rendering and update on every time step.  The way I've understood to do this is instead decouple those.  Update should take a 'time passed since last frame' to allow for varying frame rates, and one way to do draw / render is to lock it to an FPS (which you can intentionally slow down if things get too complex).  The way it's coded in the tutorial, each person's computer may process frames at a different speed so the 'speed' variables for moving entities are different for everyone. 
